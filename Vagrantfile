# ENV['VAGRANT_DEFAULT_PROVIDER'] = 'libvirt'
Vagrant.configure("2") do |config|
    config.vm.define "Luma-Workshop" do |vm|
        vm.vm.box = "generic/ubuntu2204"
        vm.vm.hostname = "Luma-Workshop"
        vm.vm.provider :libvirt do |vb|
            vb.cpus=6
            vb.memory=8000
            vb.default_prefix=""
        end
    end
end