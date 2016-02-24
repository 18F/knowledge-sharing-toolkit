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
  pages-sites
  lunr-server
  team-api
)

DATA_CONTAINERS = %w(
  pages-sites
)


DAEMON_TO_DATA_CONTAINERS = {
  'nginx-18f' => ['pages-sites:ro'],
  'pages' => ['pages-sites'],
}

def_command :build_images, 'Build Docker images' do |args|
  (args.empty? ? IMAGES : args).each do |image|
    message = "Building #{image}"
    marker = '-' * message.size
    puts "#{marker}\n#{message}\n#{marker}"
    exec_cmd "docker build -t #{image} -f ./#{image}/Dockerfile ./#{image}"
  end
end

def_command :create_data_containers, 'Create Docker data containers' do |args|
  (args.empty? ? DATA_CONTAINERS : args).each do |container_name|
    exec_cmd "docker run --name #{container_name} #{container_name} " \
      "echo Created data container \\\"#{container_name}\\\""
  end
end

def _config_dir_volume_binding(image_name)
  local_config_dir = File.join(LOCAL_ROOT_DIR, image_name, 'config')
  image_config_dir = "#{APP_SYS_ROOT}/#{image_name}"
  "-v #{local_config_dir}:#{image_config_dir}:ro"
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
    "#{_volumes_from(data_containers)} #{image_name} #{command}"
end

def_command :run_daemons, 'Run Docker containers as daemons' do |args|
  (args.empty? ? IMAGES : args).each do |image|
    _run_container(image, '-d',
      data_containers: DAEMON_TO_DATA_CONTAINERS[image] || [])
  end
end

def_command :run_container, 'Run a shell within a Docker container' do |args|
  if args.size == 1
    _run_container(args.first, '-it', command: '/bin/bash -l',
      data_containers: DAEMON_TO_DATA_CONTAINERS[args.first] || [])
  else
    puts 'run_container accepts only a single container name as an argument'
  end
end

def_command :stop_daemons, 'Stop Docker containers running as daemons' do |args|
  (args.empty? ? IMAGES : args).each do |image|
    exec_cmd "if $(docker ps -a | grep -q ' #{image_name}$'); then " \
      "docker stop #{image_name}; fi"
  end
end

def_command :rm_containers, 'Remove stopped containers' do
  exec_cmd 'docker rm $(docker ps -a | sed -e \'s/.* \([^ ]*$\)/\1/\' | ' \
    'grep -v NAMES)'
end

execute_command ARGV
