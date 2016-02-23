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

APP_SYS_ROOT = '/usr/local/18f'

IMAGES = %w(
  dev-base
  dev-standard
  nginx-18f
  oauth2_proxy
  hmacproxy
  authdelegate
  18f-pages
  lunr-server
  team-api
)

def message_marker(message)
  marker = '-' * message.size
  "#{marker}\n#{message}\n#{marker}"
end

def build_image(image)
  puts message_marker("Building #{image}")
  exec_cmd "docker build -t #{image} -f ./#{image}/Dockerfile ./#{image}"
end

def_command :build_images, 'Build Docker images' do |args|
  (args.empty? ? IMAGES : args).each { |image| build_image(image) }
end

def _remove_container(image_name)
  exec_cmd "if $(docker ps -a | grep -q ' #{image_name}$'); then " \
    "docker rm #{image_name}; fi"
end

def _run_container(image_name, options, command: '')
  puts "Running: #{image_name}"

  # Remove any existing containers matching the image name.
  _remove_container(image_name)

  # Mount the corresponding config directories as volumes. Name the container
  # the same as the image.
  local_config_dir = File.join(LOCAL_ROOT_DIR, image_name, 'config')
  image_config_dir = "#{APP_SYS_ROOT}/#{image_name}"
  exec_cmd "docker run #{options} --name #{image_name} " \
    "-v #{local_config_dir}:#{image_config_dir} " \
    "#{image_name} #{command}"
end

def_command :run_daemons, 'Run Docker containers as daemons' do |args|
  args.each { |image| _run_container(image, '-d') }
end

def_command :run_container, 'Run a shell within a Docker container' do |args|
  if args.size == 1
    _run_container(args.first, '-it', command: '/bin/bash -l')
  else
    puts 'run_container accepts only a single container name as an argument'
  end
end

def_command :rm_containers, 'Remove stopped containers' do
  exec_cmd 'docker rm $(docker ps -a | sed -e \'s/.* \([^ ]*$\)/\1/\' | ' \
    'grep -v NAMES)'
end

execute_command ARGV
