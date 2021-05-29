<template>
  <div class="content">
    <div class="container-fluid">
      <div class="row">
        <div class="col-3">
          <select v-model="folders.folderId" class="form-control">
            <option v-for="folder in folders.folderOptions" :key="folder.id">
              {{ itemValue(folder) }}
            </option>
          </select>
        </div>
        <div class="col-4">
          <button id="newFolder" class="btn btn-primary" @click="newFolder">
            New Folder
          </button>
        </div>
      </div>

      <b-modal id="folder-modal" hide-footer>
        <template #modal-title> <b>New Folder</b></template>
        <div class="row mt-3">
          <label for="nameInput">Name: * </label>
          <input
            type="text"
            class="form-control"
            v-model="folders.folderName"
            id="nameInput"
          />
        </div>

        <div class="row d-flex justify-content-between">
          <div class="col-5">
            <div class="row mt-3">
              <label for="sourceInput">Source:* </label>
              <select
                type="text"
                class="form-control"
                v-model="folders.source"
                id="sourceInput"
              >
                <option
                  v-for="lang in folders.languageOptions"
                  :key="lang.id"
                  :value="lang.id"
                >
                  {{ itemValue(lang) }}
                </option>
              </select>
            </div>
          </div>
          <div class="col-5">
            <div class="row mt-3">
              <label for="destinationInput">Destination:* </label>
              <select
                type="text"
                class="form-control"
                v-model="folders.destination"
                id="destinationInput"
              >
                <option
                  v-for="lang in folders.languageOptions"
                  :value="lang.id"
                  :key="lang.id"
                >
                  {{ itemValue(lang) }}
                </option>
              </select>
            </div>
          </div>
        </div>

        <div class="row mt-3">
          <b-button class="btn btn-primary" block @click.prevent="save"
            >Save</b-button
          >
        </div>
      </b-modal>
    </div>
  </div>
</template>

<<script>
import LTable from "src/components/Table.vue";
import Card from "src/components/Cards/Card.vue";
import axios from "axios";
import URL from 'src/util/url';
export default {
  components: {
    LTable,
    Card,
  },
  data() {
    return {
      folders: {
        folderId: 0,
        folderOptions: [],
        languageOptions: [],
        folderName: "",
        source: "",
        destination: "",
        BASE_URL: URL.LOCAL_BASE
      },
    };
  },
  mounted() {
    this.getAll();
  },
  methods: {
    itemValue(folder) {
        return folder["name"];
    },
    newFolder() {
        this.folders["folderId"] = null;
        this.$bvModal.show("folder-modal");
    },
    getAll() {
      debugger
      axios.get(`${this.folders.BASE_URL}/folders-and-languages`).then((response) => {
        this.folders["folderOptions"] = response.data.folders;
        this.folders["languageOptions"] = response.data.languages;

      });
    },
    add() {
      this.$bvModal.show("folder-modal");
    },
    save() {
      let id = this.folders["folderId"];
      if (id) {
        this.edit(id);
      } else {
        this.doSave();
      }
    },
    doSave() {
        
      let name = this.folders["folderName"];
      let source = this.folders["source"];
      let destination = this.folders["destination"];

      axios
        .post(`${this.folders.BASE_URL}/folders`, {
          name,
          source :{
            id: source
          },
          destination: {
            id: destination
          },
        })
        .then(() => {
          this.$bvModal.hide("folder-modal");
          this.getAll();
        });
    },
    edit(id) {
      let value = this.types["itemName"];
      axios
        .put(`${this.folders.BASE_URL}/types/${id}`, {
          value: value,
          id: id,
        })
        .then(() => {
          this.$bvModal.hide("type-modal");
          this.getAll();
        });
    },
    showEdit(item) {
      this.types["folderName"] = item.name;
      this.types["folderId"] = item.id;
      this.$bvModal.show("type-modal");
    },
    deleteItem(item) {
      axios.delete(`${this.folders.BASE_URL}/types/${item.id}`).then(() => {
        this.getAll();
      });
    },
  },
};
</script>