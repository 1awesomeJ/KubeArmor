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
## Containerd:
I have also implemented an OCI hook in a containerd runtime.
Below are the steps I followed:

1. I stopped and deleted my existing minikube node.
   ```minikube stop```
   
   ```minikube delete```
3. I started minikube using containerd as the container runtime
```minikube start  --container-runtime=containerd ```

4. I logged into the minikube VM using ```minikube ssh```

5. Inside the minikube VM, I created three files ```hooklog.txt```,   ```base-spec.json``` and  ```prestarthook.sh```
 ```
   sudo touch /etc/containerd/base-spec.json
   sudo touch /usr/local/share/hooklog.txt
   sudo vi /usr/local/bin/prestarthook.sh
```


```prestarthook.sh```  has the following content:

```
docker@minikube:~$ cat /usr/local/bin/prestarthook.sh 
#!/bin/bash

echo "Setting default container runtime for pods..."
export DEFAULT_CONTAINER_RUNTIME="containerd"
echo "This prestart hook was successfully executed at $(date)" >> /usr/local/share/hooklog.txt
docker@minikube:~$ 

```

I created the content of ```base-spec.json``` like so:

```
ctr oci spec >> spec.txt
sudo cp spec.txt /etc/containerd/base-spec.json
```
Then I added the following using my text editor:
```
"hooks": {
      "createContainer": [
         {
            "path": "/usr/local/bin/prestarthook.sh",
            "args": []
         }
      ],
      "poststop": [
         {
            "path": "/usr/local/bin/prestarthook.sh",
            "args": []
         }
      ]
  }

```

```hooklog.txt``` was an ampty file.

5. I set the right permissions on the three files:

```
sudo chmod 666 /etc/containerd/base-spec.json
sudo chmod 666 /usr/local/share/hooklog.txt 
sudo chmod 777 /usr/local/bin/prestarthook.sh

```
6. I edited the ```/etc/containerd/config.toml``` and set the ```base_runtime_spec``` like so:
```
 [plugins."io.containerd.grpc.v1.cri".containerd.runtimes]
        [plugins."io.containerd.grpc.v1.cri".containerd.runtimes.runc]
          runtime_type = "io.containerd.runc.v2"
```
```
          base_runtime_spec = "/etc/containerd/base-spec.json"
```
```
          [plugins."io.containerd.grpc.v1.cri".containerd.runtimes.runc.options]
            SystemdCgroup = true
```
7. I restarted my containerd service:
   ```sudo systemctl restart containerd```
8. I exited the ssh connection to the minikube VM

9. I created a new pod
``` kubectl run test-pod --image=docker.io/nginx:latest```

10. I logged back into the minkube VM
```minikube ssh```

11. I checked the content of ```hooklog.txt```, and it had the expected content like so:
```
docker@minikube:~$ cat /usr/local/share/hooklog.txt
This prestart hook was successfully executed at Mon Feb 19 11:07:13 UTC 2024
This prestart hook was successfully executed at Mon Feb 19 11:28:13 UTC 2024

```
I also plan on adding the steps for implmenting OCI hooks with a docker runtime to the design document soon.

Thank you!!!
