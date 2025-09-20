export interface Product {
  id: number;
  created_at: Date;
  updated_at: Date;
  name: string;
  description: string;
  price: number;
  vendor_id: number;
  remaining: number;
  category_id: number;
}
