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
      <NumberField source="vendor_id" />
      <NumberField source="remaining" />
      <NumberField source="price" />
      <ReferenceField source="category_id" reference="categories">
        <TextField source="name" />
      </ReferenceField>
    </SimpleShowLayout>
  </Show>
);

export default ProductShow;
