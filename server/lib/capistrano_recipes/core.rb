# set :writeable_group, "http"
# set :shared_dirs, []
# set :writeable_dirs, []
# set :symlinks, []

Capistrano::Configuration.instance.load do
  namespace :core do
    desc "[internal] Initialize variables"
    task :init, :except => { :no_release => true } do
      set :shared_dirs, fetch(:shared_dirs, []) + %w[public/uploads db config]
      set :writeable_dirs, fetch(:writeable_dirs, []) + %w[log pids public/uploads db]
      set :symlinks, fetch(:symlinks, []) + %w[public/uploads config/database.yml] 
    end

    desc "[internal] Change directory group permission and setup database.yml"
    task :setup, :except => { :no_release => true } do
      dirs = shared_dirs.map {|d| File.join(shared_path, d)}
      run "#{try_sudo} mkdir -p #{dirs.join(' ')}"
      run "#{try_sudo} chmod g+w #{dirs.join(' ')}" if fetch(:group_writable, true)

      dirs = writeable_dirs.map { |d| File.join(shared_path, d) }
      run "#{try_sudo} chgrp #{writeable_group} #{dirs.join(' ')}"

      config = ERB.new(File.read("config/database.yml.erb"))
      put config.result(binding), "#{shared_path}/config/database.yml"
    end

    desc "[internal] Creates the symlink to uploads shared folder for the most recently deployed version."
    task :symlink, :except => { :no_release => true } do
      cmds = []
      symlinks.each { |name|
        source = File.exists?("#{shared_path}/#{name}") ? name : File.basename(name)
        cmds << "rm -rf #{release_path}/#{name} 2>/dev/null; ln -sfn #{shared_path}/#{source} #{release_path}/#{name}"
      }

      run cmds.join("; ")
    end

    on :start, "core:init"
    after "deploy:setup",           "core:setup"
    after "deploy:finalize_update", "core:symlink"
  end
end
