sudo iptables-restore << EOT
# Generated by iptables-save v1.4.4 on Tue Nov  1 21:01:44 2011
*filter
:INPUT DROP [2029:131259]
:FORWARD DROP [0:0]
:OUTPUT ACCEPT [1649:176446]
-A INPUT -i lo -j ACCEPT 
-A INPUT -m state --state RELATED,ESTABLISHED -j ACCEPT 
-A INPUT -p tcp -m tcp --dport 22 -j ACCEPT 
-A INPUT -p tcp -m tcp --dport 443 -j ACCEPT 
-A OUTPUT -o lo -j ACCEPT 
COMMIT
# Completed on Tue Nov  1 21:01:44 2011
EOT