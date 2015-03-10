Capistrano::Configuration.instance.load do
  namespace :deploy do
    task :start, :roles => :app, :except => { :no_release => true } do
      run "systemctl start #{application}"
    end
   
    task :stop, :roles => :app, :except => { :no_release => true } do
      run "systemctl stop #{application}"
    end
   
    task :restart, :roles => :app, :except => { :no_release => true } do
      run "systemctl restart #{application}"
    end
  end
end
