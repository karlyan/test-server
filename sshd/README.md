
## sample run command

```bash
docker run \
-v $(pwd)/key.pub:/root/.ssh/authorized_keys \
-v $(pwd)/sshdConfig:/etc/ssh \
--name="sshd-2025" --network host --restart always -d karlyan/sshd:2025.01
```

```bash
ssh -i key.private -p 2233 root@<IP>
```

