import { List, Datagrid, TextField, NumberField, DateField } from "react-admin";

const ProductList = () => (
  <List>
    <Datagrid>
      <TextField source="id" />
      <TextField source="name" />
      <TextField source="description" />
      <DateField source="created_at" showTime />
      <DateField source="updated_at" showTime />
      <NumberField source="price" />
      <NumberField source="remaining" />
      <NumberField source="vendor_id" />
      <NumberField source="category_id" />
    </Datagrid>
  </List>
);

export default ProductList;
