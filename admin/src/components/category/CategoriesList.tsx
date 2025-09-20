import { DataTable, List } from 'react-admin';

const CategoryList = () => (
    <List>
        <DataTable>
            <DataTable.Col source="id" />
            <DataTable.Col source="name" />
        </DataTable>
    </List>
);

export default CategoryList;
