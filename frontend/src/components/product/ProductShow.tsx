import {
  Show,
  SimpleShowLayout,
  TextField,
  DateField,
  ReferenceField,
  NumberField,
} from "react-admin";

const ProductShow = () => (
  <Show>
    <SimpleShowLayout>
      <TextField source="id" />
      <DateField source="created_at" />
      <DateField source="updated_at" />
      <TextField source="name" />
      <TextField source="description" />
      <ReferenceField source="vendor_id" reference="users">
        <TextField source="name" />
      </ReferenceField>
      <NumberField source="remaining" />
      <NumberField source="price" />
      <ReferenceField source="category_id" reference="categories" />
    </SimpleShowLayout>
  </Show>
);

export default ProductShow;
