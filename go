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

command_group :build, 'Image building commands'

LOCAL_ROOT_DIR = File.absolute_path(File.dirname(__FILE__))
APP_SYS_ROOT = '/usr/local/18f'

IMAGES = %w(
  dev-base
  dev-standard
  nginx-18f
  oauth2_proxy
  hmacproxy
  authdelegate
  pages
  lunr-server
  team-api
)

DATA_CONTAINERS = {
  'pages-data' => 'pages',
  'team-api-data' => 'team-api',
}

DAEMON_TO_DATA_CONTAINERS = {
  'lunr-server' => ['pages-data:ro'],
  'nginx-18f' => [
    'pages-data:ro',
    'team-api-data:ro',
  ],
  'pages' => ['pages-data:rw'],
  'oauth2_proxy' => [],
  'hmacproxy' => [],
  'authdelegate' => [],
  'team-api' => ['team-api-data:rw'],
}

NEEDS_SSH = %w(team-api)

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
  return known_containers if args.empty?
  _check_names(args, known_containers, 'data container')
end

def _daemons(args)
  daemons = DAEMON_TO_DATA_CONTAINERS.keys
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

def_command :create_data_containers, 'Create Docker data containers' do |args|
  _data_containers(args).each do |container_name|
    base_image = DATA_CONTAINERS[container_name]
    exec_cmd "if ! $(docker ps -a | grep -q ' #{container_name}$'); then " \
      "docker run --name #{container_name} #{base_image} " \
      "echo Created data container \\\"#{container_name}\\\" " \
      "from \\\"#{base_image}\\\"; fi"
  end
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
    "#{_volumes_from(DATA_CONTAINERS.keys)} #{image_name} #{command}"
end

def_command :run_daemons, 'Run Docker containers as daemons' do |args|
  _daemons(args).each do |image|
    _run_container(image, '-d',
      data_containers: DAEMON_TO_DATA_CONTAINERS[image])
  end
end

def_command :run_container, 'Run a shell within a Docker container' do |args|
  if args.empty?
    puts 'run_container accepts a container name and an argument list'
  end
  image = args.shift
  _images([image])
  command = args.empty? ? '/bin/bash' : args.join(' ')
  _run_container(image, '-it', command: command,
    data_containers: DAEMON_TO_DATA_CONTAINERS[args.first])
end

def_command :stop_daemons, 'Stop Docker containers running as daemons' do |args|
  _daemons(args).each do |image|
    exec_cmd "if $(docker ps -a | grep -q ' #{image_name}$'); then " \
      "docker stop #{image_name}; fi"
  end
end

def_command :rm_containers, 'Remove stopped (non-data) containers' do |args|
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

execute_command ARGV
