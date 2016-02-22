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

IMAGES= %w(
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

def build_image(image)
  message = "Building #{image}"
  marker = '-' * message.size
  puts "#{marker}\n#{image}\n#{marker}"
  exec_cmd "docker build -t #{image} -f ./#{image}/Dockerfile ./#{image}"
end

def_command :build_images, 'Build Docker images' do |args|
  (args.empty? ? IMAGES : args).each { |image| build_image(image) }
end

execute_command ARGV
