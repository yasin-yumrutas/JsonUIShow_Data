import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

interface Item {
  id: number;
  name: string;
  price: number;
}

@Component({
  selector: 'app-items',
  templateUrl: './items.component.html',
  styleUrls: ['./items.component.css']
})
export class ItemsComponent implements OnInit {
  items: Item[] = [];

  constructor(private http: HttpClient) {}

  ngOnInit(): void {
    this.fetchItems();
  }

  fetchItems(): void {
    this.http.get<Item[]>('http://localhost:8080/api/items').subscribe({
      next: (items) => {
        this.items = items;
      },
      error: (err) => {
        console.error('An error occurred:', err);
      }
    });
  }
}






------> Bu örnekte, Angular ile bu API servisine nasıl istek göndereceğinizi gösteren temel bir örnek veriyorum. Angular projenizde uygun bileşenleri oluşturmanız gerekecektir.

------>Bu Angular bileşeni, HttpClient kullanarak Go API servisinden öğeleri alır ve bileşen içinde görüntüler.

Lütfen kodu ihtiyacınıza göre uyarlayarak kullanın. API yolu, port numarası ve diğer bağlantı detayları ihtiyaca göre düzenlenmelidir.
