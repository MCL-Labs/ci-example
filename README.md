# Gitlab CI Example Project

This is an example project to demonstrate how to use Gitlab CI.


## Install Gitlab Runner

- the Runner server can be installed on the same server as GitLab, or on a different server. 
- Runner Server and Gitlab server must be able to communicate with each other.

```bash
# Download the binary for your system
sudo curl -L --output /usr/local/bin/gitlab-runner https://gitlab-runner-downloads.s3.amazonaws.com/latest/binaries/gitlab-runner-linux-amd64

# Give it permission to execute
sudo chmod +x /usr/local/bin/gitlab-runner

# Create a GitLab Runner user
sudo useradd --comment 'GitLab Runner' --create-home gitlab-runner --shell /bin/bash

# Install and run as a service
sudo gitlab-runner install --user=gitlab-runner --working-directory=/home/gitlab-runner
sudo gitlab-runner start


# Register the Runner
gitlab-runner register  --url http://gitlab.landui.cn  --token glrt-4ugeFzsf2WcnRwtkeWNf

# Runtime platform                                    arch=amd64 os=linux pid=354049 revision=b92ee590 version=17.4.0
# WARNING: Running in user-mode.
# WARNING: The user-mode requires you to manually start builds processing:
# WARNING: $ gitlab-runner run
# WARNING: Use sudo for system-mode:
# WARNING: $ sudo gitlab-runner...

# Enter the GitLab instance URL (for example, https://gitlab.com/):
# [http://gitlab.landui.cn]: https://gitlab.landui.cn/
# Verifying runner... is valid                        runner=4ugeFzsf2
# Enter a name for the runner. This is stored only in the local config.toml file:
# [MathxH]: MathxH
# Enter an executor: custom, ssh, docker+machine, kubernetes, docker-autoscaler, instance, shell, parallels, virtualbox, docker, docker-windows:
# shell
# Runner registered successfully. Feel free to start it, but if it's running already the config should be automatically reloaded!

# Configuration (with the authentication token) was saved in "/home/mathxh/.gitlab-runner/config.toml"

# Start the Runner
gitlab-runner start

# Check the status
gitlab-runner status

# Stop the Runner
gitlab-runner stop
```