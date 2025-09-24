import { EmailField, Show, SimpleShowLayout, TextField } from "react-admin";

const UserShow = () => (
  <Show>
    <SimpleShowLayout>
      <TextField source="id" />
      <TextField source="name" />
      <EmailField source="email" />
    </SimpleShowLayout>
  </Show>
);

export default UserShow;
