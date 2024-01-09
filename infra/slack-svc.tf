resource "digitalocean_app" "captainslog-slack" {
  spec {
    name   = "captainslog-slack"
    region = "tor"
    domain {
      name = "malaspina.xyz"
      type = "PRIMARY"
      zone = "malaspina.xyz"
    }

    service {
      name               = "captainslog-slack"
      environment_slug   = "go"
      instance_count     = 1
      instance_size_slug = "professional-xs"

      github {
        repo = "elliotrushton/captainslog-slack"
        branch         = "main"
	deploy_on_push = "true"
      }
    }
  }
}
