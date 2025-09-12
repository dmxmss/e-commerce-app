import { Admin, Resource, ListGuesser } from "react-admin";
import { Layout } from "./Layout";
import dataProvider from "./dataProvider.ts";
import ProductList from "./components/product/ProductList";
import ProductShow from "./components/product/ProductShow";

const App = () => (
  <Admin layout={Layout} dataProvider={dataProvider} disableTelemetry>
    <Resource name="products" list={ProductList} show={ProductShow} />
    <Resource name="categories" list={ListGuesser} />
  </Admin>
);

export default App;
