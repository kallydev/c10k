test_concurrency:
	wrk -t12 -c2000 -d15s "http://localhost/"

install_go:
	apt install -y build-essential
	snap install go --classic

install_wrk:
	apt install -y build-essential libssl-dev git unzip
	git submodule init
	git submodule update
	cd ./tool/wrk && make
	cp ./tool/wrk/wrk /usr/local/bin

enable_bbr:
	echo net.core.default_qdisc=fq >> /etc/sysctl.conf
	echo net.ipv4.tcp_congestion_control=bbr >> /etc/sysctl.conf
	sysctl -p

enable_swap:
	fallocate -l 2G /swap
	chmod 600 /swap
	mkswap /swap
	swapon /swap
	cp /etc/fstab /etc/fstab.backup
	echo '/swap none swap sw 0 0' | tee -a /etc/fstab
