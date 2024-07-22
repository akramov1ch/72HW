# Uyga vazifa: RabbitMQ Routing Xususiyatidan Foydalanish orqali Buyurtma Boshqaruv Tizimi

## Maqsad
Ushbu vazifaning maqsadi `RabbitMQ` ning routing xususiyati bilan ishlashni o'rganish va uni murakkab tizimlarda qanday qo'llashni tushunishdir.

## Talablar
1. **REST API (Producer) yaratish**: 
    - Buyurtmalarni yaratish va oqish uchun endpointlarga ega `REST API` ni amalga oshiring.
    - Yangi buyurtma yaratilganda, buyurtma tafsilotlarini `RabbitMQ` `direct exchange` ga yuboring. Har bir buyurtma tegishli routing kaliti bilan jo'natiladi (masalan, "`order.pending`", "`order.completed`", "`order.canceled`").

2. **Worker (Consumer) yaratish:**: 
    - Har bir turdagi buyurtmalarni qayta ishlash uchun alohida consumerlar yarating.
    - Consumer `RabbitMQ` dan mos buyurtmalarni olishi va qayta ishlashi kerak, masalan:
        - Buyurtma tafsilotlarini `MongoDB` ga saqlash.
        - Buyurtma holatini yangilash.

## Qo'shimcha Talablar
- Buyurtma Qayta Ishlashni Yaxshilash:
    - Buyurtmalarni qayta ishlashdan oldin validatsiya qiling.
    - Har xil buyurtma holatlarini qo'llab-quvvatlang (masalan, ko'rib chiqilgan, bajarilgan, bekor qilingan).
