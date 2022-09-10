# Kubernetes On Raspberry PIs

## Prerequisites

* At least 2 Raspberry PIs (Model 3 and above)
* Alpine Linux installed on all the PIs.

**The following steps needs to be performed on all the PIs unless
otherwise specified.**

## Repositories

Enable edge/community & edge/testing in `/etc/apk/repositories`

## Installing required software

```bash
doas apk update
doas apk add containerd kubelet kubeadm cni-plugins cni-plugin-flannel
```

## Enable Linux cgroups memory feature

```bash
doas sed -i 's/$/ cgroup_enable=cpuset cgroup_memory=1 cgroup_enable=memory/' /boot/cmdline.txt
```

## Load Kubernetes required modules

```bash
cat <<EOF | doas tee /etc/modules-load.d/k8s.conf >/dev/null
overlay
br_netfilter
EOF
```

## Required network configurations

```bash
cat <<EOF | doas tee /etc/sysctl.d/k8s.conf >/dev/null
net.bridge.bridge-nf-call-iptables  = 1
net.bridge.bridge-nf-call-ip6tables = 1
net.ipv4.ip_forward                 = 1
EOF
```

## Install CNI plugins

```bash
curl -OL https://github.com/containernetworking/plugins/releases/download/v1.1.1/cni-plugins-linux-arm64-v1.1.1.tgz
curl -OL https://github.com/containernetworking/plugins/releases/download/v1.1.1/cni-plugins-linux-arm64-v1.1.1.tgz.sha256
sha256sum -c cni-plugins-linux-arm64-v1.1.1.tgz.sha256 && \
    doas mkdir /opt/cni/bin -p && \
    doas tar xzf cni-plugins-linux-arm64-v1.1.1.tgz -C /opt/cni/bin
```

## Enable Containerd and Kubelet

```bash
doas rc-update add containerd
doas rc-update add kubelet
```

## Reboot

## Control Plane

**Run only on control plane node**

```bash
doas kubeadm init --pod-network-cidr=10.244.0.0/16
```

* If your raspberry pi has < 2 GB of memory like me, ask kubeadm to
  ignore the minimum memory error,

```bash
doas kubeadm init --pod-network-cidr=10.244.0.0/16 --ignore-preflight-errors=Mem
```

* Get Kube Config

```bash
mkdir -p $HOME/.kube
doas cp -f /etc/kubernetes/admin.conf $HOME/.kube/config
doas chown $(id -u):$(id -g) $HOME/.kube/config
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
doas kubeadm join 192.168.0.100:6443 \
    --token XXXXXX.XXXXXXXXXXXXXXXX \
    --discovery-token-ca-cert-hash sha256:XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
```
