<template>
  <div class="content">
    <div class="container-fluid">
      <div class="row"></div>
      <div class="row">
        <div class="col-12">
          <card
            class="strpied-tabled-with-hover"
            body-classes="table-full-width table-responsive"
          >
            <template slot="header">
              <div class="row">
                <div class="col-11">
                  <h4 class="card-title">Languages</h4>
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
                :columns="languages.columns"
                :data="languages.data"
                v-on:editItem="showEdit"
                v-on:deleteItem="deleteItem"
              >
              </l-table>
            </div>
          </card>
        </div>
      </div>
    </div>

    <b-modal id="language-modal"  content-class="header-class-modal" title="New Language" hide-backdrop>
      <template>
        <label class="font-weight-bold" for="nameInput">Name:* </label>
        <input
          type="text"
          class="form-control"
          v-model="languages.itemName"
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
const tableColumns = ["Name"];
export default {
  components: {
    LTable,
    Card,
  },
  data() {
    return {
      languages: {
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
      axios.get(`${this.languages.BASE_URL}/languages`).then((response) => {
        this.languages["data"] = response.data;
      });
    },
    add() {
       this.languages["itemName"] = ""
      this.$bvModal.show("language-modal");
    },
    save() {
      let id = this.languages["itemId"];
      if (id) {
        this.edit(id);
      } else {
        this.doSave();
      }
    },
    doSave() {
      let name = this.languages["itemName"];
      axios
        .post(`${this.languages.BASE_URL}/languages`, {
          name,
        })
        .then(() => {
          this.$bvModal.hide("language-modal");
          this.getAll();
        });
    },
    edit(id) {
      let name = this.languages["itemName"];
      axios
        .put(`${this.languages.BASE_URL}/languages/${id}`, {
          name: name,
          id: id,
        })
        .then(() => {
          this.$bvModal.hide("language-modal");
          this.getAll();
        });
    },
    showEdit(item) {
      this.languages["itemName"] = item.name;
      this.languages["itemId"] = item.id;
      this.$bvModal.show("language-modal");
    },
    deleteItem(item) {
      debugger;
      axios
        .delete(`${this.languages.BASE_URL}/languages/${item.id}`)
        .then(() => {
          this.getAll();
        });
    },
  },
};
</script>
<style>
</style>
