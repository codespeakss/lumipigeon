scp ~/codespeakss/lumipigeon/web/* 67_root:/var/www/html/;
ssh 67_root "pkill lumipigeon" ; scp ~/codespeakss/lumipigeon/out/lumipigeon 67_root:/root ; ssh 67_root " /root/lumipigeon"