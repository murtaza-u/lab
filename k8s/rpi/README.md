# Kubernetes On Raspberry PIs

## Prerequisites

* At least 2 Raspberry PIs (Model 3 and above)
* [Raspberry Pi OS
  (64-bit)](https://downloads.raspberrypi.org/raspios_lite_arm64/images/raspios_lite_arm64-2022-04-07/2022-04-04-raspios-bullseye-arm64-lite.img.xz)
  installed on all the PIs.

**The following steps needs to be performed on all the PIs unless
otherwise specified.**

## Installing required software

```bash
sudo apt install -y software-properties-common apt-transport-https \
    ca-certificates curl gnupg lsb-release
```

* Containerd

```bash
# Add docker repository
curl -fsSL https://download.docker.com/linux/debian/gpg | \
    gpg --dearmor -o /usr/share/keyrings/docker.gpg

cat <<EOF | sudo tee /etc/apt/sources.list.d/docker.list >/dev/null
deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker.gpg] https://download.docker.com/linux/debian $(lsb_release -cs) stable
EOF

sudo apt update -y
sudo apt install -y containerd.io
```

* kubeadm, kubelet and kubectl

```bash
curl -fsSLo /usr/share/keyrings/kubernetes-archive-keyring.gpg \
    https://packages.cloud.google.com/apt/doc/apt-key.gpg

cat <<EOF | sudo tee /etc/apt/sources.list.d/kubernetes.list >/dev/null
deb [signed-by=/usr/share/keyrings/kubernetes-archive-keyring.gpg] https://apt.kubernetes.io/ kubernetes-xenial main
EOF

sudo apt update -y
sudo apt install -y kubelet kubeadm kubectl

# prevent automatic updates
sudo apt-mark hold kubelet kubeadm kubectl
```

## Enable Linux cgroups memory feature

```bash
sudo sed -i 's/$/ cgroup_enable=cpuset cgroup_memory=1 cgroup_enable=memory/' /boot/cmdline.txt
```

## Disable swap

```bash
sudo dphys-swapfile swapoff
sudo dphys-swapfile uninstall
sudo apt purge -y dphys-swapfile
```

## Disable Bluetooth(optional)

```bash
cat <<EOF | sudo tee /etc/modprobe.d/raspi-blacklist.conf >/dev/null
# Bluetooth
blacklist btbcm
blacklist hci_uart
EOF

sudo systemctl disable bluetooth.service
```

## Load Kubernetes required modules

```bash
cat <<EOF | sudo tee /etc/modules-load.d/k8s.conf >/dev/null
overlay
br_netfilter
EOF
```

## Required network configurations

```bash
cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf >/dev/null
net.bridge.bridge-nf-call-iptables  = 1
net.bridge.bridge-nf-call-ip6tables = 1
net.ipv4.ip_forward                 = 1
EOF
```

## Containerd configuration

```bash
sudo mkdir /etc/containerd -p
containerd config default | sudo tee /etc/containerd/config.toml >/dev/null
```

* To use the systemd cgroup driver in `/etc/containerd/config.toml` with
`runc`, set

```bash
[plugins."io.containerd.grpc.v1.cri".containerd.runtimes.runc]
  ...
  [plugins."io.containerd.grpc.v1.cri".containerd.runtimes.runc.options]
    SystemdCgroup = true
```

## Reboot

## Control Plane

**Run only on control plane node**

```bash
sudo kubeadm init --pod-network-cidr=10.244.0.0/16
```

* If your raspberry pi has < 2 GB of memory like me, ask kubeadm to
  ignore the minimum memory error,

```bash
sudo kubeadm init --pod-network-cidr=10.244.0.0/16 --ignore-preflight-errors=Mem
```

* Get Kube Config

```bash
mkdir -p $HOME/.kube
sudo cp -f /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config
```

* Install `flannel`

> Flannel is a network fabric for containers, designed for Kubernetes.
> It is Responsible for allocating a subnet lease to each host out of a
> larger, preconfigured address space.

```bash
kubectl apply -f https://raw.githubusercontent.com/flannel-io/flannel/master/Documentation/kube-flannel.yml
```

## Worker Nodes

* Once all the all pods are ready, ask worker nodes to join the
  cluster.

```bash
~$ kubectl get pods --all-namespaces
NAMESPACE      NAME                                    READY   STATUS    RESTARTS      AGE
kube-flannel   kube-flannel-ds-2mb4g                   1/1     Running   3 (85m ago)   5h10m
kube-flannel   kube-flannel-ds-fkf9p                   1/1     Running   1 (94m ago)   125m
kube-flannel   kube-flannel-ds-vsq4c                   1/1     Running   3             5h22m
kube-system    coredns-6d4b75cb6d-7r55q                1/1     Running   3 (94m ago)   5h24m
kube-system    coredns-6d4b75cb6d-ll27f                1/1     Running   3 (94m ago)   5h24m
kube-system    etcd-control-plane                      1/1     Running   3 (94m ago)   5h18m
kube-system    kube-apiserver-control-plane            1/1     Running   3 (94m ago)   5h18m
kube-system    kube-controller-manager-control-plane   1/1     Running   3 (94m ago)   5h19m
kube-system    kube-proxy-cnrcv                        1/1     Running   1 (94m ago)   125m
kube-system    kube-proxy-tddj4                        1/1     Running   2 (85m ago)   5h10m
kube-system    kube-proxy-vcg98                        1/1     Running   3 (94m ago)   5h24m
kube-system    kube-scheduler-control-plane            1/1     Running   3 (94m ago)   5h18m
```

**Run only on worker nodes**

```bash
sudo kubeadm join 192.168.0.100:6443 \
    --token XXXXXX.XXXXXXXXXXXXXXXX \
    --discovery-token-ca-cert-hash sha256:XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```

```bash
kubeadm token create --print-join-command
```
