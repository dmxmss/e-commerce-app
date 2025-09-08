import { Admin, ListGuesser, Resource } from "react-admin";
import { Layout } from "./Layout";
import simpleRestProvider from "ra-data-simple-rest";

const App = () => (
  <Admin layout={Layout} dataProvider={simpleRestProvider("http://localhost:3000/api")} disableTelemetry>
    <Resource name="products" list={ListGuesser} />
    <Resource name="users" list={ListGuesser} />
  </Admin>
);

export default App;
