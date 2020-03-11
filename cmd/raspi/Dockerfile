FROM python:3.7.4-alpine
RUN apk add docker bash git jq curl vim go nodejs python2 python3
RUN apk add linux-headers g++ python2-dev python3-dev

SHELL ["/bin/bash", "-c"]

RUN pip3 install --upgrade pip
RUN pip3 install requests websocket-client schedule pyyaml psutil rpi.gpio gpiozero
# RUN pip3 install awscli

# RUN curl -o /tmp/tmp https://cli-assets.heroku.com/branches/stable/heroku-linux-arm.tar.gz
# RUN tar zxvf /tmp/tmp -C / && rm -f /tmp/tmp
# RUN ln -s /heroku/bin/heroku /bin

# RUN curl -o /tmp/tmp https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-sdk-245.0.0-linux-x86_64.tar.gz
# RUN tar zxvf /tmp/tmp -C / &&  rm -f /tmp/tmp
# RUN /google-cloud-sdk/install.sh --quiet
# RUN ln -s /google-cloud-sdk/bin/{gcloud,gsutil,bq} /bin

# RUN curl -o /tmp/tmp https://releases.hashicorp.com/terraform/0.12.21/terraform_0.12.21_linux_arm.zip
# RUN unzip -d /bin /tmp/tmp && rm -f /tmp/tmp
