import ipaddress

# Define the VPC CIDR and allocated subnets
vpc_cidr = ""
allocated = [
    "",
]

# Create the VPC network object
vpc = ipaddress.IPv4Network(vpc_cidr)
allocated_networks = [ipaddress.IPv4Network(cidr) for cidr in allocated]

# Start with the entire VPC CIDR and subtract allocated networks
remaining = [vpc]
for network in allocated_networks:
    new_remaining = []
    for r in remaining:
        # Only exclude if the allocated network is within the current range
        if network.overlaps(r):
            new_remaining.extend(r.address_exclude(network))
        else:
            new_remaining.append(r)  # Keep the current range as-is
    remaining = new_remaining

# Print the available CIDR ranges
print("Available CIDR ranges:")
for r in remaining:
    print(r)
