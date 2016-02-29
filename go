#! /usr/bin/env ruby

require 'English'

Dir.chdir File.dirname(__FILE__)

def try_command_and_restart(command)
  exit $CHILD_STATUS.exitstatus unless system command
  exec({ 'RUBYOPT' => nil }, RbConfig.ruby, *[$PROGRAM_NAME].concat(ARGV))
end

begin
  require 'bundler/setup' if File.exist? 'Gemfile'
rescue LoadError
  try_command_and_restart 'gem install bundler'
rescue SystemExit
  try_command_and_restart 'bundle install'
end

begin
  require 'go_script'
rescue LoadError
  try_command_and_restart 'gem install go_script' unless File.exist? 'Gemfile'
  abort "Please add \"gem 'go_script'\" to your Gemfile"
end

extend GoScript
check_ruby_version '2.3.0'

command_group :build, 'Image and container building commands'

LOCAL_ROOT_DIR = File.absolute_path(File.dirname(__FILE__))
APP_SYS_ROOT = '/usr/local/18f'
NETWORK = '18f/knowledge-sharing-toolkit'
GITHUB_REPOSITORY = '18F/knowledge-sharing-toolkit.git'

IMAGES = %w(
  dev-base
  dev-standard
  oauth2_proxy
  hmacproxy
  authdelegate
  pages
  lunr-server
  team-api
  nginx
)

DATA_CONTAINERS = {
  'pages-data' => 'pages',
  'team-api-data' => 'team-api',
}

DAEMONS = {
  'lunr-server' => {
    data_containers: ['pages-data:ro'],
  },
  'pages' => {
    data_containers: ['pages-data:rw']
  },
  'oauth2_proxy' => {
    data_containers: [],
  },
  'hmacproxy' => {
    data_containers: [],
  },
  'authdelegate' => {
    data_containers: [],
  },
  'team-api' => {
    data_containers: ['team-api-data:rw'],
  },
  'nginx' => {
    flags: '-p 80:80 -p 443:443',
    data_containers: [
      'pages-data:ro',
      'team-api-data:ro',
    ],
  },
}

NEEDS_SSH = %w(team-api)

REMOTE_HOST = 'ubuntu@hub.18f.gov'
REMOTE_ROOT = 'knowledge-sharing-toolkit'
SECRETS_BUNDLE_NAME = '18f-knowledge-sharing-toolkit-secrets'
SECRETS_BUNDLE_FILE = "#{SECRETS_BUNDLE_NAME}.tar.bz2"
SECRET_FILES = %w(
  */config/env-secret.sh
  nginx/config/auth/pages-passwords.txt
  nginx/config/ssl/dhparam*
  nginx/config/ssl/keys/*
  pages/config/pages.secret
  ssh/config/id_rsa*
  ssh/config/known_hosts*
)

def _check_names(names, collection, type_label)
  names.each do |name|
    next if collection.include?(name)
    puts "\"#{name}\" does not match any known #{type_label}; " \
      "valid #{type_label}s are:\n  #{collection.join("\n  ")}"
    exit 1
  end
  names
end

def _images(args)
  args.empty? ? IMAGES : _check_names(args, IMAGES, 'image')
end

def _data_containers(args)
  known_containers = DATA_CONTAINERS.keys
  args.empty? ?
    known_containers :
    _check_names(args, known_containers, 'data container')
end

def _daemons(args)
  daemons = DAEMONS.keys
  args.empty? ? daemons : _check_names(args, daemons, 'daemon')
end

def_command :build_images, 'Build Docker images' do |args|
  _images(args).each do |image|
    message = "Building #{image}"
    marker = '-' * message.size
    puts "#{marker}\n#{message}\n#{marker}"
    exec_cmd "docker build -t #{image} -f ./#{image}/Dockerfile ./#{image}"
  end
end

def_command :create_data_containers, 'Create data containers' do |args|
  _data_containers(args).each do |container_name|
    base_image = DATA_CONTAINERS[container_name]
    exec_cmd "if ! $(docker ps -a | grep -q ' #{container_name}$'); then " \
      "docker run --name #{container_name} #{base_image} " \
      "echo Created data container \\\"#{container_name}\\\" " \
      "from \\\"#{base_image}\\\"; fi"
  end
end

command_group :run_containers, 'Container running commands'

def _network_is_running
  `docker network ls`.split("\n")[1..-1]
    .map { |network| network.gsub(/  */, ' ').split[1] }
    .include?(NETWORK)
end

def_command :create_network, 'Start the local network between containers' do
  if !_network_is_running
    exec_cmd "docker network create --driver bridge #{NETWORK}"
  end
end

def_command :rm_network, 'Start the local network between containers' do
  exec_cmd "docker network rm #{NETWORK}" if _network_is_running
end

def _config_dir_volume_binding(image_name)
  local_config_dir = File.join(LOCAL_ROOT_DIR, image_name, 'config')
  image_config_dir = "#{APP_SYS_ROOT}/#{image_name}/config"
  "-v #{local_config_dir}:#{image_config_dir}:ro"
end

def _ssh_config_dir_volume_binding(image_name)
  NEEDS_SSH.include?(image_name) ? _config_dir_volume_binding('ssh') : ''
end

def _volumes_from(data_containers)
  data_containers.map { |container| "--volumes-from #{container}" }.join(' ')
end

def _run_container(image_name, options, command: '', data_containers: [])
  puts "Running: #{image_name}"

  # Remove any existing containers matching the image name.
  exec_cmd "if $(docker ps -a | grep -q ' #{image_name}$'); then " \
    "docker rm #{image_name}; fi"

  # Mount the corresponding config directories as volumes. Name the container
  # the same as the image.
  exec_cmd "docker run #{options} --name #{image_name} " \
    "#{_config_dir_volume_binding(image_name)} " \
    "#{_ssh_config_dir_volume_binding(image_name)} " \
    "#{_network_is_running ? "--net=#{NETWORK}" : '' } " \
    "#{_volumes_from(data_containers)} #{image_name} #{command}"
end

def_command :run_daemons, 'Run Docker containers as daemons' do |args|
  _daemons(args).each do |daemon_name|
    daemon = DAEMONS[daemon_name]
    _run_container(daemon_name, "-d #{daemon[:flags]}",
      data_containers: daemon[:data_containers])
  end
end

def_command :run_container, 'Run a shell within a Docker container' do |args|
  if args.empty?
    puts 'run_container accepts a container name and an argument list'
  end
  image = args.shift
  _images([image])
  command = args.empty? ? '/bin/bash' : args.join(' ')
  data_containers = (DAEMONS[image] || {})[:data_containers]
  _run_container(image, '-it', command: command,
    data_containers: data_containers || [])
end

def_command :reload_nginx, 'Reload Nginx after a config change' do
  exec_cmd 'docker kill -s HUP nginx'
end

def_command :run_hmacproxy, 'Run hmacproxy that will sign requests' do |args|
  if args.size != 1
    puts "You must specify a single upstream host as an argument to " \
      "run_hmacproxy."
    exit 1
  end
  upstream_host = args.first
  exec_cmd "docker run -d --name hmacproxy-sign --net=#{NETWORK} " \
    "-p 8084:8084 #{_config_dir_volume_binding('hmacproxy')} " \
    "hmacproxy run-proxy #{upstream_host}"
  puts "Requests to http://#{`docker-machine ip`.rstrip}:8084 " \
    "will be forwarded to #{upstream_host}."
end

def_command :stop_hmacproxy, 'Stop the hmacproxy signing container' do
  exec_cmd 'if $(docker ps | grep -q \'hmacproxy-sign$\'); then ' \
    'docker stop hmacproxy-sign; docker rm hmacproxy-sign; fi'
end

def_command :stop_daemons, 'Stop containers running as daemons' do |args|
  _daemons(args).each do |daemon|
    exec_cmd "if $(docker ps -a | grep -q ' #{daemon}$'); then " \
      "docker stop #{daemon}; fi"
  end
end

command_group :cleanup, 'Image and container cleanup commands'

def_command :rm_containers, 'Remove stopped non-data containers' do |args|
  images = _images(args)
  containers = `docker ps -a`.split("\n")[1..-1]
    .map { |container| container.match(/ ([^ ]*)$/)[1] }
    .reject { |container| container.end_with?('-data') }
    .select { |container| images.include?(container) }
  exec_cmd "docker rm #{containers.join(' ')}" unless containers.empty?
end

def_command :rm_images, 'Remove unused images' do
  unused_images = `docker images`.split("\n")[1..-1]
    .select { |image| image.start_with?('<none>') }
    .map { |image| image.gsub(/  */, ' ').split(' ')[2] }
  exec_cmd "docker rmi #{unused_images.join(' ')}" unless unused_images.empty?
end

command_group :remote, 'Remote server access and management'

def _exec_remote(remote_command)
  exec_cmd "ssh #{REMOTE_HOST} 'cd #{REMOTE_ROOT} && #{remote_command}'"
end

def_command :init_remote, 'Initialize the remote system repository' do
  exec_cmd "ssh #{REMOTE_HOST} " \
    "'git clone git@github.com:#{GITHUB_REPOSITORY} #{REMOTE_ROOT}'"
end

def_command :sync_remote, 'Synchronize the remote system repository' do
  _exec_remote 'git fetch origin master && ' \
    'git clean -f && git reset --hard origin/master'
end

def_command :ssh_remote, 'Open an interactive SSH session to the remote' do
  exec "ssh -t #{REMOTE_HOST} 'cd #{REMOTE_ROOT} && exec /bin/bash -l'"
end

command_group :secrets, 'Commands to manage system secrets'

def _ensure_secrets_bundle_does_not_exist
  if File.exist?(SECRETS_BUNDLE_FILE)
    puts "Secret bundle file #{SECRETS_BUNDLE_FILE} already exists;\n" \
      "please delete or rename it before running this command."
    exit 1
  end
end

def _ensure_secrets_bundle_exists
  if !File.exist?(SECRETS_BUNDLE_FILE)
    puts "Secret bundle file #{SECRETS_BUNDLE_FILE} does not exist;\n" \
      "please run `./go bundle_secrets` before running this command."
    exit 1
  end
end

def_command :bundle_secrets, 'Create a bundle from local secret files' do
  _ensure_secrets_bundle_does_not_exist
  secret_files = Dir.glob(SECRET_FILES)
  exec_cmd "tar cvf #{SECRETS_BUNDLE_NAME}.tar #{secret_files.join(' ')}"
  exec_cmd "bzip2 -9 #{SECRETS_BUNDLE_NAME}.tar"
end

def_command :unpack_secret_bundle, 'Unpack the secret bundle' do
  _ensure_secret_bundle_exists
  exec_cmd "bzip2 -dc #{SECRETS_BUNDLE_FILE} | tar xvf -"
end

def_command :push_secrets, 'Push the secret bundle and unpack it' do
  bundle_secrets if !File.exist?(SECRETS_BUNDLE_FILE)
  exec_cmd "scp #{SECRETS_BUNDLE_FILE} #{REMOTE_HOST}:#{REMOTE_ROOT}/"
  _exec_remote "bzip2 -dc #{SECRETS_BUNDLE_FILE} | tar xvf -"
end

def_command :fetch_secrets, 'Fetch the secret bundle from the remote host' do
  _ensure_secrets_bundle_does_not_exist
  _exec_remote "rm -f #{SECRETS_BUNDLE_FILE} && ruby ./go bundle_secrets"
  exec_cmd "scp #{REMOTE_HOST}:#{REMOTE_ROOT}/#{SECRETS_BUNDLE_FILE} ."
end

command_group :system, 'Commands to start and stop the entire system'

def_command :start, 'Start the entire system' do
  puts "Starting the system...\nCreating network #{NETWORK}:"
  create_network
  puts 'Creating data containers (if they don\'t already exist):'
  create_data_containers
  puts 'Running daemon containers:'
  run_daemons
  puts 'System start complete.'
end

def_command :stop, 'Stop the entire system' do
  puts "Stopping the system...\nStopping all daemons:"
  stop_daemons
  puts 'Removing non-data containers:'
  rm_containers
  puts "Stopping network #{NETWORK}:"
  rm_network
  puts 'System stop complete.'
end

execute_command ARGV
