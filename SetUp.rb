#!/usr/bin/env ruby

dir = File.expand_path(__dir__)
MY_BASH = File.join(dir, '.jonathon_bash_profile')
MY_BIN = File.join(dir, 'bin')
MY_VIM = File.join(dir, '.jonathon_vimrc')
MY_GIT_C = File.join(dir, '.gitconfig')

home_dir = File.expand_path('~')
LOCAL_BASH = File.join(home_dir, '.bashrc')
LOCAL_VIM = File.join(home_dir, '.vimrc')
LOCAL_GIT_C = File.join(home_dir, '.gitconfig')

def exist_and_doesnt_already_contain(my, local)
  return puts("File not found #{my}") unless File.exist?(my)
  return puts("File not found #{local}") unless File.exist?(local)

  return puts("Skipping #{my}...") if File.foreach(local).any? { |l| l.include?(my) }

  open(local, 'a') { |f| yield(f) }
end

def can_i?(message)
  puts "#{message}?"
  ['yes', 'y'].include?(gets.chomp.downcase)
end

# bash rc
exist_and_doesnt_already_contain(MY_BASH, LOCAL_BASH) do |file|
  next unless can_i?("Source bash rc")
  puts("Sourcing #{MY_BASH}...")
  file << "source #{MY_BASH}"
end
# bin
exist_and_doesnt_already_contain(MY_BIN, LOCAL_BASH) do |file|
  next unless can_i?("Export bin")
  puts("Exporing path #{MY_BIN}...")
  file << "export PATH=\"$PATH:#{MY_BIN}\""
end
# vim rc
exist_and_doesnt_already_contain(MY_VIM, LOCAL_VIM) do |file|
  next unless can_i?("Source vmrc")
  puts("Sourcing #{MY_VIM}...")
  file << "source #{MY_VIM}"
end
# git config
exist_and_doesnt_already_contain(MY_GIT_C, LOCAL_GIT_C) do |file|
  next unless can_i?("Include gitconfig")
  puts("Including #{MY_GIT_C}...")
  file << "  path = #{MY_GIT_C}"
end
