import { DataTable, EmailField, List } from 'react-admin';

const UserList = () => (
  <List>
    <DataTable>
      <DataTable.Col source="id" />
      <DataTable.Col source="name" />
      <DataTable.Col source="email">
        <EmailField source="email" />
      </DataTable.Col>
    </DataTable>
  </List>
);

export default UserList;
