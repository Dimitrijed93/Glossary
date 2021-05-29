<template>
  <div class="content">
    <div class="container-fluid">
      <div class="row"></div>
      <div class="row">
        <div class="col-12">
          <card class="strpied-tabled-with-hover"
            body-classes="table-full-width table-responsive">
            <template slot="header">
              <div class="row">
                <div class="col-11">
                  <h4 class="card-title">Types</h4>
                </div>
                <div class="col-1">
                  <button
                    type="submit"
                    class="btn btn-info btn-fill float-left"
                    @click.prevent="add"
                  >
                    Add
                  </button>
                </div>
              </div>
            </template>

            <div class="table-responsive">
              <l-table
                class="table-hover"
                :columns="types.columns"
                :data="types.data"
                v-on:editItem="showEdit"
                v-on:deleteItem="deleteItem"
              >
              </l-table>
            </div>
          </card>
        </div>
      </div>
    </div>

    <b-modal id="type-modal" title="New Type" hide-backdrop>
      <template>
        <label class="font-weight-bold" for="nameInput">Name: * </label>
        <input
          type="text"
          class="form-control"
          v-model="types.itemName"
          id="nameInput"
        />
      </template>
      <template #modal-footer>
        <b-button class="btn btn-fill btn-info" block @click.prevent="save"
          >Save</b-button
        ></template
      >
    </b-modal>
  </div>
</template>

<script>
import LTable from "src/components/Table.vue";
import Card from "src/components/Cards/Card.vue";
import axios from "axios";
const tableColumns = ["Value"];
export default {
  components: {
    LTable,
    Card,
  },
  data() {
    return {
      types: {
        columns: [...tableColumns],
        data: [],
        itemName: "",
        itemId: "",
        BASE_URL: URL.LOCAL_BASE,
      },
    };
  },
  mounted() {
    this.getAll();
  },
  methods: {
    getAll() {
      axios.get(`${this.types.BASE_URL}/types`).then((response) => {
        this.types["data"] = response.data;
      });
    },
    add() {
      this.$bvModal.show("type-modal");
    },
    save() {
      let id = this.types["itemId"];
      if (id) {
        this.edit(id);
      } else {
        this.doSave();
      }
    },
    doSave() {
      let value = this.types["itemName"];
      axios
        .post(`${this.types.BASE_URL}/types`, {
          value,
        })
        .then(() => {
          this.$bvModal.hide("type-modal");
          this.getAll();
        });
    },
    edit(id) {
      let value = this.types["itemName"];
      axios
        .put(`${this.types.BASE_URL}/types/${id}`, {
          value: value,
          id: id,
        })
        .then(() => {
          this.$bvModal.hide("type-modal");
          this.getAll();
        });
    },
    showEdit(item) {
      this.types["itemName"] = item.value;
      this.types["itemId"] = item.id;
      this.$bvModal.show("type-modal");
    },
    deleteItem(item) {
      axios.delete(`${this.types.BASE_URL}/types/${item.id}`).then(() => {
        this.getAll();
      });
    },
  },
};
</script>
<style>
</style>
