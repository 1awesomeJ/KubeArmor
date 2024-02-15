Dear Mentors, I have SUCCESSFULLY implemented an OCI hook that runs with CRIO container runtime!

I am using an Ubuntu VM
Here are the steps I followed:

1. I installed the docker packages 
```sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin```

2. I installed kubectl 
```
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
```

3. I installed minikube
```
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube
```
4. I started minikube using crio as the container runtime
```minikube start  --container-runtime=crio ```

5. I logged into the minikube VM using ```minikube ssh```

6. Inside the minikube VM, I created three files ```hook_log.txt```,   ```kubearmor-hook.json``` and  ```prestarthook.sh``` inside the  ```/usr/share/containers/oci/hooks.d/``` directory.

```prestarthook.sh```  has the following content:

```
#!/bin/bash

echo "Setting the desired default container runtime to 'crio'"
export DEFAULT_CONTAINER_RUNTIME="crio"
echo "OCI hook executed successfully! at $(date)" >> /usr/share/containers/oci/hooks.d/hook_log.txt
```

```kubearmor-hook.json``` has the following content:

```
{
    "version": "1.0.0",
    "hook": {
        "path": "/usr/share/containers/oci/hooks.d/prestarthook.sh"
    },
    "when": {
        "always": true 
    },
    "stages": ["createContainer", "poststop"] 
}
```

```hook_log.txt``` was an ampty file.

8. I set the right permissions on ```kubearmor-hook.json``` and ```hook_log.txt```

```
sudo chmod a+x /usr/local/bin/prestarthook.sh
sudo chmod 777 /usr/share/containers/oci/hooks.d/hook_log.txt
```

9. I exited the ssh connection to the minikube VM

10. I created a new pod
``` kubectl run test-pod --image=docker.io/nginx:latest```

11. I logged back into the minkube VM
```minikube ssh```

12. I checked the content of ```hook_log.txt```, and it had the expected content like so:
```
docker@minikube:~$ cat /usr/share/containers/oci/hooks.d/hook_log.txt
OCI hook executed successfully! at Wed Feb 14 15:31:54 UTC 2024
docker@minikube:~$ 
```

In the meantime, I will work on creating and testing an OCI hook with containerd runtime.
I also plan on adding the steps for implmenting OCI hooks with a docker runtime to the design document soon.

Thank you!!!
