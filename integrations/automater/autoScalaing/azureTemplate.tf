resource "azurerm_virtual_machine_scale_set" "example" {
    name                = "example"
      location            = "eastus"
        resource_group_name = azurerm_resource_group.example.name
          upgrade_policy_mode = "Automatic"

          sku {
                name     = "Standard_DS1_v2"
                    capacity = 2
                      
          }

          storage_profile_image_reference {
                publisher = "Canonical"
                    offer     = "UbuntuServer"
                        sku       = "18.04-LTS"
                            version   = "latest"
                              
          }

          storage_profile_os_disk {
                caching              = "ReadWrite"
                    create_option        = "FromImage"
                        managed_disk_type    = "Standard_LRS"
                            disk_size_gb         = "30"
                              
          }

          os_profile {
                computer_name_prefix = "example"
                    admin_username       = "example"
                        admin_password       = "example"
                          
          }

          os_profile_linux_config {
                disable_password_authentication = false
                  
          }

          network_profile {
                name    = "example"
                    primary = true

                    ip_configuration {
                            name      = "example"
                                  primary   = true
                                        subnet_id = azurerm_subnet.example.id
                                            
                    }
                      
          }
          
}

