#!/usr/bin/env bash

set -e

instance_type=t2.micro
ami_id=ami-0f5ee92e2d63afc18
vpc_id="$(aws ec2 describe-vpcs --filters 'Name=isDefault,Values=true' --query 'Vpcs[0].VpcId' --output text)"
subnet_id="$(aws ec2 describe-subnets --filters "Name=vpc-id,Values=$vpc_id" --query "Subnets[0].SubnetId" --output text)"
instance_id="$(aws ec2 run-instances \
  --image-id "$ami_id" \
  --subnet-id "$subnet_id" \
  --instance-type "$instance_type" \
  --query "Instances[0].InstanceId" \
  --output text)"

echo "waiting for $instance_id to start..."
aws ec2 wait instance-running --instance-ids "$instance_id"

echo "$instance_id is up and running"
read -r -p "press [Enter] to terminate instance"

aws ec2 terminate-instances --instance-ids "$instance_id" >/dev/null
echo "terminating instance $instance_id..."
aws ec2 wait instance-terminated --instance-ids "$instance_id"
echo "$instance_id terminated"