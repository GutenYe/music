$: << File.expand_path("../../lib", __FILE__) 
require "bundler/capistrano"
require "capistrano_recipes/core"
require "capistrano_recipes/systemd"

set :application, "milks"
set :repository,  "/home/guten/dev/one/music/server"
set :scm, :git
set :deploy_to, "/apps/#{application}"
set :user, "root"
set :writeable_group, "http"
set :keep_releases, 1
set :server_url, "http://api.milk.priv"

server "localhost", :app, :web, :db, :primary => true

after "deploy:restart", "deploy:cleanup"

task :update_server_url do
  run "echo '#{server_url}' > #{release_path}/SERVER_URL"
end

after "deploy:finalize_update", "update_server_url"
