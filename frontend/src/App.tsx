import { Admin, Resource, ListGuesser } from "react-admin";
import { Layout } from "./Layout";
import dataProvider from "./dataProvider.ts";
import ProductList from "./components/product/ProductList";
import ProductShow from "./components/product/ProductShow";
import CategoriesList from "./components/categories/CategoriesList.tsx";

const App = () => (
  <Admin layout={Layout} dataProvider={dataProvider} disableTelemetry>
    <Resource name="products" list={ProductList} show={ProductShow} />
    <Resource name="categories" list={CategoriesList} />
  </Admin>
);

export default App;
