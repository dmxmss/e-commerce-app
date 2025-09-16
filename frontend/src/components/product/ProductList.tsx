import { List, Datagrid, TextField, NumberField, DateField, ReferenceField } from "react-admin";

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
      <ReferenceField source="vendor_id" reference="users" sortBy="name">
        <TextField source="name" />
      </ReferenceField>
      <ReferenceField source="category_id" reference="categories" sortBy="name">
        <TextField source="name" />
      </ReferenceField>
    </Datagrid>
  </List>
);

export default ProductList;
