import { DataProvider } from "react-admin";

const apiUrl = "http://localhost:8080/api";

const dataProvider: DataProvider = {
  getList: async (resource, params) => {
    let request = `${apiUrl}/${resource}?`;

    if (params.pagination.page && params.pagination.perPage) {
      request += `page=${params.pagination.page}&perPage=${params.pagination.perPage}&`;
    }

    if (params.sort.field && params.sort.order) {
      request += `sortField=${params.sort.field}&sortOrder=${params.sort.order}&`;
    }

    const response = await fetch(request);
    const json = await response.json();
    return {
      data: json.data,
      total: json.total,
    };
  },

  getOne: async (resource, params) => {
    console.log("resource: ", resource, " params: ", params);
    const request = `${apiUrl}/${resource}/${params.id}`;

    const response = await fetch(request);
    const data = await response.json();

    console.log("data:", data);
    return {
      data: data,
    };
  },

  getMany: async (resource, params) => {
    let ids = "";
    for (id of params.ids) {
      ids += `ids=${id}&`;
    }

    const request = `${apiUrl}/${resource}?${ids}`;

    const response = await fetch(request);
    const data = await response.json();

    return {
      data: data,
    };
  },

  getManyReference: async (resource, params) => {
    let request = `${apiUrl}/${resource}?${params.target}=${params.id}&`;

    if (params.pagination.page && params.pagination.perPage) {
      request += `page=${params.pagination.page}&perPage=${params.pagination.perPage}&`;
    }

    if (params.sort.field && params.sort.order) {
      request += `sortField=${params.sort.field}&sortOrder=${params.sort.order}&`;
    }

    const response = await fetch(request);
    const data = await response.json();

    return {
      data: data,
      total: data.length,
    };
  },

  create: async (resource, params) => {
    const request = `${apiUrl}/${resource}`;

    const response = await fetch(request, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(params.data),
    });

    const data = await response.json();
    return { data };
  },

  update: async (resource, params) => {
    const request = `${apiUrl}/${resource}/${params.id}`;

    const response = await fetch(request, {
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(params.data),
    });

    const data = await response.json();
    return { data };
  },

  updateMany: async (resource, params) => {
    const responses = await Promise.All(
      params.ids.map((id) => {
        fetch(`${apiUrl}/${resource}/${id}`, {
          method: "PUT",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify(params.data),
        }).then((res) => res.json());
      }),
    );
    return { data: responses.map((r) => r.id) };
  },

  delete: async (resource, params) => {
    await fetch(`${apiUrl}/${resource}/${params.id}`, {
      method: "DELETE",
    });

    return { data: { id: params.id } };
  },

  deleteMany: async (resource, params) => {
    await Promise.All(
      params.ids.map((id) => {
        fetch(`${apiUrl}/${resource}/${id}`, { method: "DELETE" });
      }),
    );
    return { data: params.ids };
  },
};

export default dataProvider;
