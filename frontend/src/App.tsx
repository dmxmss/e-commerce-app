import { Admin, Resource, ListGuesser, ShowGuesser } from "react-admin";
import { Layout } from "./Layout";
import dataProvider from "./dataProvider.ts";

import ProductList from "./components/product/ProductList";
import ProductShow from "./components/product/ProductShow";

import CategoriesList from "./components/category/CategoriesList";

import UserList from "./components/user/UserList"

const App = () => (
  <Admin layout={Layout} dataProvider={dataProvider} disableTelemetry>
    <Resource name="products" list={ProductList} show={ProductShow} />
    <Resource name="categories" list={CategoriesList} />
    <Resource name="users" list={UserList} show={ShowGuesser} />
  </Admin>
);

export default App;
