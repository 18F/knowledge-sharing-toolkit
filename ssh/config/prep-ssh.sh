_SSH_CONFIG_DIR=$APP_SYS_ROOT/ssh/config

if [ ! -d $HOME/.ssh ]; then
  mkdir -m 700 $HOME/.ssh

  for f in id_rsa id_rsa.pub known_hosts
  do
    cp $_SSH_CONFIG_DIR/$f $HOME/.ssh/$f
  done
  chmod 600 $HOME/.ssh/*
fi

eval $(ssh-agent -s)
ssh-add
