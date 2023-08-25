exports.ids = [8,3];
exports.modules = {

/***/ 38:
/***/ (function(module, exports, __webpack_require__) {

// style-loader: Adds some css to the DOM by adding a <style> tag

// load the styles
var content = __webpack_require__(44);
if(content.__esModule) content = content.default;
if(typeof content === 'string') content = [[module.i, content, '']];
if(content.locals) module.exports = content.locals;
// add CSS to SSR context
var add = __webpack_require__(4).default
module.exports.__inject__ = function (context) {
  add("74c952df", content, true, context)
};

/***/ }),

/***/ 43:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_vue_style_loader_index_js_ref_3_oneOf_1_0_node_modules_css_loader_dist_cjs_js_ref_3_oneOf_1_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_ref_3_oneOf_1_2_node_modules_nuxt_components_dist_loader_js_ref_0_0_node_modules_vue_loader_lib_index_js_vue_loader_options_DogVaccines_vue_vue_type_style_index_0_id_45b42ebf_prod_lang_css___WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(38);
/* harmony import */ var _node_modules_vue_style_loader_index_js_ref_3_oneOf_1_0_node_modules_css_loader_dist_cjs_js_ref_3_oneOf_1_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_ref_3_oneOf_1_2_node_modules_nuxt_components_dist_loader_js_ref_0_0_node_modules_vue_loader_lib_index_js_vue_loader_options_DogVaccines_vue_vue_type_style_index_0_id_45b42ebf_prod_lang_css___WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(_node_modules_vue_style_loader_index_js_ref_3_oneOf_1_0_node_modules_css_loader_dist_cjs_js_ref_3_oneOf_1_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_ref_3_oneOf_1_2_node_modules_nuxt_components_dist_loader_js_ref_0_0_node_modules_vue_loader_lib_index_js_vue_loader_options_DogVaccines_vue_vue_type_style_index_0_id_45b42ebf_prod_lang_css___WEBPACK_IMPORTED_MODULE_0__);
/* harmony reexport (unknown) */ for(var __WEBPACK_IMPORT_KEY__ in _node_modules_vue_style_loader_index_js_ref_3_oneOf_1_0_node_modules_css_loader_dist_cjs_js_ref_3_oneOf_1_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_ref_3_oneOf_1_2_node_modules_nuxt_components_dist_loader_js_ref_0_0_node_modules_vue_loader_lib_index_js_vue_loader_options_DogVaccines_vue_vue_type_style_index_0_id_45b42ebf_prod_lang_css___WEBPACK_IMPORTED_MODULE_0__) if(["default"].indexOf(__WEBPACK_IMPORT_KEY__) < 0) (function(key) { __webpack_require__.d(__webpack_exports__, key, function() { return _node_modules_vue_style_loader_index_js_ref_3_oneOf_1_0_node_modules_css_loader_dist_cjs_js_ref_3_oneOf_1_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_ref_3_oneOf_1_2_node_modules_nuxt_components_dist_loader_js_ref_0_0_node_modules_vue_loader_lib_index_js_vue_loader_options_DogVaccines_vue_vue_type_style_index_0_id_45b42ebf_prod_lang_css___WEBPACK_IMPORTED_MODULE_0__[key]; }) }(__WEBPACK_IMPORT_KEY__));


/***/ }),

/***/ 44:
/***/ (function(module, exports, __webpack_require__) {

// Imports
var ___CSS_LOADER_API_IMPORT___ = __webpack_require__(3);
var ___CSS_LOADER_EXPORT___ = ___CSS_LOADER_API_IMPORT___(function(i){return i[1]});
// Module
___CSS_LOADER_EXPORT___.push([module.i, "#vaccine-list{height:400px;overflow-y:scroll}#popup{height:64%;overflow-y:scroll}", ""]);
// Exports
___CSS_LOADER_EXPORT___.locals = {};
module.exports = ___CSS_LOADER_EXPORT___;


/***/ }),

/***/ 47:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
// ESM COMPAT FLAG
__webpack_require__.r(__webpack_exports__);

// CONCATENATED MODULE: ./node_modules/babel-loader/lib??ref--2-0!./node_modules/vue-loader/lib/loaders/templateLoader.js??ref--6!./node_modules/@nuxt/components/dist/loader.js??ref--0-0!./node_modules/vue-loader/lib??vue-loader-options!./components/DogVaccines.vue?vue&type=template&id=45b42ebf&
var render = function render() {
  var _vm = this,
    _c = _vm._self._c;
  return _c('div', {
    staticClass: "card"
  }, [_vm._ssrNode("<div class=\"card-content\"><div id=\"vaccine-list\" class=\"content\"><table class=\"table is-hoverable is-fullwidth\"><thead id=\"header-table\"><tr><th>Nombre</th> <th>Fecha</th></tr></thead> <tbody id=\"table-body\">" + _vm._ssrList(_vm.vaccines, function (vaccine, index) {
    return "<tr><td>" + _vm._ssrEscape(_vm._s(vaccine.NombreVacuna)) + "</td> <td>" + _vm._ssrEscape(_vm._s(vaccine.Fecha)) + "</td></tr>";
  }) + "</tbody></table></div></div> "), _vm._ssrNode("<div" + _vm._ssrClass("modal", {
    'is-active': _vm.showModalFlag
  }) + ">", "</div>", [_vm._ssrNode("<div class=\"modal-background\"></div> "), _vm._ssrNode("<div id=\"popup\" class=\"modal-card\">", "</div>", [_vm._ssrNode("<header class=\"modal-card-head\"><p class=\"modal-card-title\">Agregar una Vacuna</p> <button aria-label=\"close\" class=\"delete\"></button></header> "), _vm._ssrNode("<section class=\"modal-card-body\">", "</section>", [_vm._ssrNode("<div class=\"field\"><label class=\"label\">Nombre</label> <div class=\"control\"><input type=\"text\" placeholder=\"Text input\"" + _vm._ssrAttr("value", _vm.vaccineName) + " class=\"input\"></div></div> "), _vm._ssrNode("<div class=\"field\">", "</div>", [_vm._ssrNode("<label class=\"label\">Fecha</label> "), _vm._ssrNode("<div class=\"control\">", "</div>", [_c('client-only', [_c('date-picker', {
    attrs: {
      "placeholder": "MM/DD/YYYY",
      "format": "MM/dd/yyyy"
    },
    model: {
      value: _vm.date_today,
      callback: function ($$v) {
        _vm.date_today = $$v;
      },
      expression: "date_today"
    }
  })], 1)], 1)], 2)], 2), _vm._ssrNode(" <footer class=\"modal-card-foot\"><button class=\"button is-success\">Save changes</button> <button class=\"button\">Cancel</button></footer>")], 2)], 2), _vm._ssrNode(" <footer class=\"card-footer\"><a href=\"#\" data-target=\"modal-js-example\" class=\"card-footer-item js-modal-trigger\">Agregar una vacuna</a></footer>")], 2);
};
var staticRenderFns = [];

// CONCATENATED MODULE: ./components/DogVaccines.vue?vue&type=template&id=45b42ebf&

// CONCATENATED MODULE: ./node_modules/babel-loader/lib??ref--2-0!./node_modules/@nuxt/components/dist/loader.js??ref--0-0!./node_modules/vue-loader/lib??vue-loader-options!./components/DogVaccines.vue?vue&type=script&lang=js&
/* harmony default export */ var DogVaccinesvue_type_script_lang_js_ = ({
  props: ['idDog'],
  data() {
    return {
      showModalFlag: false,
      okPressed: false,
      vaccineName: '',
      date: '',
      date_today: new Date(),
      vaccines: []
    };
  },
  methods: {
    async getVaccinesByDog(id) {
      const path = "http://localhost:8000" + '/perro/vacunas/' + id.toString();
      try {
        const res = await this.$axios.get(path);
        const data = res.data;
        this.vaccines = data;
      } catch (error) {
        console.log(error);
      }
    },
    async registerVaccine() {
      const date = new Date(this.date);
      const data = {
        NombreVacuna: this.vaccineName,
        Fecha: this.date_today,
        Perro_id: parseInt(this.idDog)
      };
      const path = "http://localhost:8000" + '/vacuna/';
      try {
        const res = await this.$axios.post(path, data);
      } catch (error) {
        console.log(error);
      }
    },
    rechargeList() {
      this.registerVaccine();
      this.getVaccinesByDog(this.idDog);
    },
    showModal() {
      this.okPressed = false;
      this.showModalFlag = true;
    },
    okModal() {
      this.okPressed = true;
      this.showModalFlag = false;
    },
    cancelModal() {
      this.okPressed = false;
      this.showModalFlag = false;
    }
  },
  async created() {
    await this.getVaccinesByDog(this.idDog);
  }
});
// CONCATENATED MODULE: ./components/DogVaccines.vue?vue&type=script&lang=js&
 /* harmony default export */ var components_DogVaccinesvue_type_script_lang_js_ = (DogVaccinesvue_type_script_lang_js_); 
// EXTERNAL MODULE: ./node_modules/vue-loader/lib/runtime/componentNormalizer.js
var componentNormalizer = __webpack_require__(2);

// CONCATENATED MODULE: ./components/DogVaccines.vue



function injectStyles (context) {
  
  var style0 = __webpack_require__(43)
if (style0.__inject__) style0.__inject__(context)

}

/* normalize component */

var component = Object(componentNormalizer["a" /* default */])(
  components_DogVaccinesvue_type_script_lang_js_,
  render,
  staticRenderFns,
  false,
  injectStyles,
  null,
  "685997fa"
  
)

/* harmony default export */ var DogVaccines = __webpack_exports__["default"] = (component.exports);

/* nuxt-component-imports */
installComponents(component, {Footer: __webpack_require__(6).default})


/***/ }),

/***/ 57:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
// ESM COMPAT FLAG
__webpack_require__.r(__webpack_exports__);

// CONCATENATED MODULE: ./node_modules/babel-loader/lib??ref--2-0!./node_modules/vue-loader/lib/loaders/templateLoader.js??ref--6!./node_modules/@nuxt/components/dist/loader.js??ref--0-0!./node_modules/vue-loader/lib??vue-loader-options!./pages/perros/_id.vue?vue&type=template&id=2d85a994&
var render = function render() {
  var _vm = this,
    _c = _vm._self._c;
  return _c('div', {
    staticClass: "tile is-ancestor"
  }, [_vm._ssrNode("<div class=\"tile is-child box\"><p class=\"title\">Datos del Perro</p> <p class=\"subtitle\">" + _vm._ssrEscape("Nombre: " + _vm._s(_vm.dog.Nombre)) + "</p> <p>" + _vm._ssrEscape("Raza: " + _vm._s(_vm.dog.Raza)) + "</p> <p>" + _vm._ssrEscape("Color: " + _vm._s(_vm.dog.Color)) + "</p> <p>" + _vm._ssrEscape("Edad: " + _vm._s(_vm.dog.Edad)) + "</p></div> <div class=\"tile is-child box\"><p class=\"title\">Datos del Dueño</p> <p class=\"subtitle\">" + _vm._ssrEscape("Nombre: " + _vm._s(_vm.owner.Nombre)) + "</p> <p>" + _vm._ssrEscape("Edad: " + _vm._s(_vm.owner.Edad)) + "</p> <p>" + _vm._ssrEscape("Sexo: " + _vm._s(_vm.owner.Sexo)) + "</p></div> "), _vm._ssrNode("<div class=\"tile is-child\">", "</div>", [_vm._ssrNode("<p class=\"title\">Vacunas</p> "), _c('DogVaccines', {
    attrs: {
      "idDog": _vm.$route.params.id
    }
  })], 2)], 2);
};
var staticRenderFns = [];

// CONCATENATED MODULE: ./pages/perros/_id.vue?vue&type=template&id=2d85a994&

// EXTERNAL MODULE: ./components/DogVaccines.vue + 4 modules
var DogVaccines = __webpack_require__(47);

// EXTERNAL MODULE: external "vuex"
var external_vuex_ = __webpack_require__(35);

// CONCATENATED MODULE: ./node_modules/babel-loader/lib??ref--2-0!./node_modules/@nuxt/components/dist/loader.js??ref--0-0!./node_modules/vue-loader/lib??vue-loader-options!./pages/perros/_id.vue?vue&type=script&lang=js&


/* harmony default export */ var _idvue_type_script_lang_js_ = ({
  data() {
    return {
      dog: {
        Nombre: '',
        Raza: '',
        Color: '',
        Edad: ''
      },
      owner: {
        Nombre: '',
        Edad: ''
      },
      vaccines: []
    };
  },
  components: {
    DogVaccines: DogVaccines["default"]
  },
  methods: {
    async getDog(id) {
      const dogID = id.toString();
      console.log(dogID);
      const path = "http://localhost:8000" + '/perro/' + dogID;
      try {
        const res = await this.$axios.get(path);
        const data = res.data;
        this.dog = data;
      } catch (error) {
        console.log(error);
      }
    },
    async getOwnerByDog(id) {
      const path = "http://localhost:8000" + '/perro/dueno/' + id.toString();
      try {
        const res = await this.$axios.get(path);
        const data = res.data;
        this.owner = data;
      } catch (error) {
        console.log(error);
      }
    },
    loadData() {
      const dogID = this.$route.params.id;
      this.getDog(dogID);
      this.getOwnerByDog(dogID);
    }
  },
  async created() {
    await this.loadData();
  }
});
// CONCATENATED MODULE: ./pages/perros/_id.vue?vue&type=script&lang=js&
 /* harmony default export */ var perros_idvue_type_script_lang_js_ = (_idvue_type_script_lang_js_); 
// EXTERNAL MODULE: ./node_modules/vue-loader/lib/runtime/componentNormalizer.js
var componentNormalizer = __webpack_require__(2);

// CONCATENATED MODULE: ./pages/perros/_id.vue





/* normalize component */

var component = Object(componentNormalizer["a" /* default */])(
  perros_idvue_type_script_lang_js_,
  render,
  staticRenderFns,
  false,
  null,
  null,
  "9dca0396"
  
)

/* harmony default export */ var _id = __webpack_exports__["default"] = (component.exports);

/* nuxt-component-imports */
installComponents(component, {DogVaccines: __webpack_require__(47).default})


/***/ })

};;
//# sourceMappingURL=_id.js.map