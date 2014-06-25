# -*- mode: ruby -*-
# vi: set ft=ruby :

# Vagrantfile API/syntax version. Don't touch unless you know what you're doing!
VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.box = "jarosser06/omnios-r151010j"
  config.vm.synced_folder ".", "/opt/zfsnap"

  config.vm.provision "shell", path: ".provision.sh"
end
