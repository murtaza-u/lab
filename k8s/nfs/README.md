# NFS

> NFSv4 supports encryption, previous versions do not. It is required
> that both client and server have NFS > v4 installed, otherwise the
> communication will fallback to previous version.

* Install required packages

```bash
sudo apk add nfs-utils
```

* Export directory

```bash
# /etc/exports
#
# See exports(5) for a description.

# use exportfs -arv to reread

# worker 1 -> 192.168.0.101
# worker 2 -> 192.168.0.102
/srv/nfs 192.168.0.101(rw,sync,secure,no_subtree_check) 192.168.0.102(rw,sync,secure,no_subtree_check)
```

* Start services (Alpine)

```bash
sudo rc-update add nfs
sudo rc-update add nfsmount

sudo rc-service nfs start
sudo rc-update add nfsmount
```

* Start services (Debian)

```bash
sudo apt install nfs-common -y
```
