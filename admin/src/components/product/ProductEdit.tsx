import {
  Edit,
  NumberInput,
  ReferenceInput,
  SimpleForm,
  TextInput,
} from "react-admin";

const ProductEdit = () => (
  <Edit>
    <SimpleForm>
      <TextInput source="id" disabled />
      <TextInput source="name" />
      <TextInput source="description" />
      <NumberInput source="remaining" />
      <NumberInput source="price" />
      <ReferenceInput source="category_id" reference="categories" />
    </SimpleForm>
  </Edit>
);

export default ProductEdit;
