# =============================================================================
# REQUIRED VARIABLES
# =============================================================================
set :application, "gophertimes"
set :domain, "www.brianketelsen.com"

set :deploy_to, "/var/www/apps/gophertimes"
set :user, "bketelsen"
set :owner, "bketelsen"
set :group, "bketelsen"
set :branch, "master"

set :use_sudo, false

set :scm, :git

set :repository,  "git@github.com:bketelsen/GopherTimes.git"
set :executable, "gophertimes"

default_run_options[:pty] = true 

# =============================================================================
# ROLES
# =============================================================================
role :web, domain
role :app, domain
role :db,  domain, :primary => true

set :default_environment, { 
  'PATH' => "/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:$PATH",
  'GOARCH'=>"386",
  'GOROOT'=>"/usr/lib/go",
  'GOPATH' => "/home/bketelsen/go",
  'GOOS'=>"linux"

}

namespace :deploy do
  task :start do 
    stop
    compile
    sudo "bluepill #{executable} start"
  end
  task :stop do 
    sudo "bluepill #{executable} stop"
  end
  task :restart do 
    stop
    start
  end

  desc "Deploy the app"
  task :default do
    update
    restart
  end

  desc "Setup and clone the repo."
  task :setup do
     # clone
    run "git clone #{repository} #{deploy_to}"
  end
  
  task :status do
    sudo "bluepill #{executable} status"
  end

  desc "compile and link the application"
  task :compile do
    clean
    run "cd #{deploy_to} && make"
  end

  desc "clean the build directory"
  task :clean do
    run "cd #{deploy_to} && make clean"
  end

  desc "Update the deployed code"
  task :update do
    run "cd #{deploy_to} && git pull origin #{branch}"
  end

  task :goinstall do
    run "goinstall -clean github.com/garyburd/twister/web"
    run "goinstall -clean github.com/garyburd/twister/server"
    run "goinstall -clean launchpad.net/mgo"
    run "goinstall -clean launchpad.net/gobson/bson"
    run "goinstall -clean github.com/knieriem/markdown"
    run "goinstall -clean github.com/garyburd/twister/expvar"
    run "goinstall -clean github.com/garyburd/twister/pprof"
  end
  
  desc "Setup and clone the repo."
  task :setup do
     # clone
    run "git clone #{repository} #{deploy_to}"
    run "goinstall github.com/garyburd/twister/web"
    run "goinstall github.com/garyburd/twister/server"
    run "goinstall launchpad.net/mgo"
    run "goinstall launchpad.net/gobson/bson"
    run "goinstall github.com/knieriem/markdown"
    run "goinstall github.com/garyburd/twister/expvar"
    run "goinstall github.com/garyburd/twister/pprof"
  end
  

end







