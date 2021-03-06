Vagrant.configure("2") do |config|

  config.vm.hostname = "api-runemadsen-com"

  config.vm.box = "opscode-ubuntu-12.04"
  config.vm.box_url = "https://opscode-vm.s3.amazonaws.com/vagrant/opscode_ubuntu-12.04_provisionerless.box"

  config.vm.network :forwarded_port, guest: 3000, host: 3001
  config.vm.network :forwarded_port, guest: 8080, host: 8081

  config.vm.provider :virtualbox do |vb|
    vb.customize ["modifyvm", :id, "--memory", "1024"]
  end

  config.omnibus.chef_version = "11.4.0"
  config.berkshelf.enabled = true

  config.vm.provision :chef_solo do |chef|
    chef.json = {
      "go" => {
        "version" => "1.1.2",
        "packages" => [
          "github.com/codegangsta/martini",
          "github.com/codegangsta/martini-contrib/...",
          "github.com/bitly/go-simplejson",
          "github.com/mattn/goreman",
          "github.com/onsi/ginkgo",
          "github.com/onsi/gomega",
          "github.com/dancannon/gorethink"
        ]
      }
    }
    chef.run_list = [
      "apt",
      "curl",
      "golang::packages",
      "rethinkdb"
    ]
  end
end