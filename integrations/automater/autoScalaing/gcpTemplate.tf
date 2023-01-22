resource "google_compute_instance_template" "example" {
    name = "example"
      machine_type = "f1-micro"
        can_ip_forward = false
          tags = ["example"]

          disk {
                source_image = "ubuntu-os-cloud/ubuntu-1804-lts"
                    auto_delete = true
                        boot = true
                          
          }

          network_interface {
                network = "default"
                  
          }
          
}

resource "google_compute_instance_group_manager" "example" {
    name = "example"
      zone = "us-central1-a"
        target_size = 2
          instance_template = google_compute_instance_template.example.self_link
          
}

