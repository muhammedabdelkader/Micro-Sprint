resource "aws_autoscaling_group" "example" {
    name                      = "example"
      max_size                  = 5
        min_size                  = 2
          desired_capacity          = 2
            launch_configuration      = aws_launch_configuration.example.name
              availability_zones        = ["us-west-2a", "us-west-2b"]
                health_check_type         = "EC2"
                tag {
                      key                 = "Name"
                          value               = "example"
                              propagate_at_launch = true
                                
                }
                
}

resource "aws_launch_configuration" "example" {
    image_id      = "ami-0c94855ba95c71c99"
      instance_type = "t2.micro"
      
}

