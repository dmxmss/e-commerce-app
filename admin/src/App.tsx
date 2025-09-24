import { Admin, Resource } from "react-admin";
import { Layout } from "./Layout";

import dataProvider from "./dataProvider.ts";
import authProvider from "./authProvider.ts";

import ProductList from "./components/product/ProductList";
import ProductShow from "./components/product/ProductShow";
import ProductEdit from "./components/product/ProductEdit";

import CategoriesList from "./components/category/CategoriesList";

import UserList from "./components/user/UserList";
import UserShow from "./components/user/UserShow";

const App = () => (
  <Admin
    layout={Layout}
    dataProvider={dataProvider}
    authProvider={authProvider}
    disableTelemetry
  >
    <Resource
      name="products"
      list={ProductList}
      show={ProductShow}
      edit={ProductEdit}
    />
    <Resource name="categories" list={CategoriesList} />
    <Resource name="users" list={UserList} show={UserShow} />
  </Admin>
);

export default App;
