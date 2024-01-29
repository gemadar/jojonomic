# jojonomic
.env tidak di hide untuk mempermudah keperluan testing

setelah melakukan pull pada repository ini, input command dibawah ini pada home directory: 

docker compose up 

Berikut endpoint beserta port:
1. Cek Harga
   :8081/api/check-harga
2. Cek Mutasi
   :8082/api/mutasi
3. Cek Saldo
   :8083/api/saldo
4. Input Harga
    :8088/api/input-harga
5. Topup
    :8085/api/topup
6. Buyback
    :8084/api/buyback

Dikarenakan tidak ada API untuk API rekening, maka telah dibuatkan satu row data rekening dengan norek : 12345678
