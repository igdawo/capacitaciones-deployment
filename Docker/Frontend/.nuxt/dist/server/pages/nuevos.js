exports.ids = [7,2];
exports.modules = {

/***/ 36:
/***/ (function(module, exports, __webpack_require__) {

// style-loader: Adds some css to the DOM by adding a <style> tag

// load the styles
var content = __webpack_require__(40);
if(content.__esModule) content = content.default;
if(typeof content === 'string') content = [[module.i, content, '']];
if(content.locals) module.exports = content.locals;
// add CSS to SSR context
var add = __webpack_require__(4).default
module.exports.__inject__ = function (context) {
  add("a9b38246", content, true, context)
};

/***/ }),

/***/ 39:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_vue_style_loader_index_js_ref_3_oneOf_1_0_node_modules_css_loader_dist_cjs_js_ref_3_oneOf_1_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_ref_3_oneOf_1_2_node_modules_nuxt_components_dist_loader_js_ref_0_0_node_modules_vue_loader_lib_index_js_vue_loader_options_DogRegister_vue_vue_type_style_index_0_id_4a236bd2_prod_lang_css___WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(36);
/* harmony import */ var _node_modules_vue_style_loader_index_js_ref_3_oneOf_1_0_node_modules_css_loader_dist_cjs_js_ref_3_oneOf_1_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_ref_3_oneOf_1_2_node_modules_nuxt_components_dist_loader_js_ref_0_0_node_modules_vue_loader_lib_index_js_vue_loader_options_DogRegister_vue_vue_type_style_index_0_id_4a236bd2_prod_lang_css___WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(_node_modules_vue_style_loader_index_js_ref_3_oneOf_1_0_node_modules_css_loader_dist_cjs_js_ref_3_oneOf_1_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_ref_3_oneOf_1_2_node_modules_nuxt_components_dist_loader_js_ref_0_0_node_modules_vue_loader_lib_index_js_vue_loader_options_DogRegister_vue_vue_type_style_index_0_id_4a236bd2_prod_lang_css___WEBPACK_IMPORTED_MODULE_0__);
/* harmony reexport (unknown) */ for(var __WEBPACK_IMPORT_KEY__ in _node_modules_vue_style_loader_index_js_ref_3_oneOf_1_0_node_modules_css_loader_dist_cjs_js_ref_3_oneOf_1_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_ref_3_oneOf_1_2_node_modules_nuxt_components_dist_loader_js_ref_0_0_node_modules_vue_loader_lib_index_js_vue_loader_options_DogRegister_vue_vue_type_style_index_0_id_4a236bd2_prod_lang_css___WEBPACK_IMPORTED_MODULE_0__) if(["default"].indexOf(__WEBPACK_IMPORT_KEY__) < 0) (function(key) { __webpack_require__.d(__webpack_exports__, key, function() { return _node_modules_vue_style_loader_index_js_ref_3_oneOf_1_0_node_modules_css_loader_dist_cjs_js_ref_3_oneOf_1_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_ref_3_oneOf_1_2_node_modules_nuxt_components_dist_loader_js_ref_0_0_node_modules_vue_loader_lib_index_js_vue_loader_options_DogRegister_vue_vue_type_style_index_0_id_4a236bd2_prod_lang_css___WEBPACK_IMPORTED_MODULE_0__[key]; }) }(__WEBPACK_IMPORT_KEY__));


/***/ }),

/***/ 40:
/***/ (function(module, exports, __webpack_require__) {

// Imports
var ___CSS_LOADER_API_IMPORT___ = __webpack_require__(3);
var ___CSS_LOADER_EXPORT___ = ___CSS_LOADER_API_IMPORT___(function(i){return i[1]});
// Module
___CSS_LOADER_EXPORT___.push([module.i, "#button{background-color:#5f9ea0;color:#f5f5f5}#button,#buttonInv{font-size:100%;font-weight:700;width:100%}#buttonInv{background-color:#f5f5f5;color:#5f9ea0}", ""]);
// Exports
___CSS_LOADER_EXPORT___.locals = {};
module.exports = ___CSS_LOADER_EXPORT___;


/***/ }),

/***/ 45:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
// ESM COMPAT FLAG
__webpack_require__.r(__webpack_exports__);

// CONCATENATED MODULE: ./node_modules/babel-loader/lib??ref--2-0!./node_modules/vue-loader/lib/loaders/templateLoader.js??ref--6!./node_modules/@nuxt/components/dist/loader.js??ref--0-0!./node_modules/vue-loader/lib??vue-loader-options!./components/DogRegister.vue?vue&type=template&id=4a236bd2&
var render = function render() {
  var _vm = this,
    _c = _vm._self._c;
  return _c('div', {
    staticClass: "box"
  }, [_vm._ssrNode("<div class=\"field\"><label class=\"label\">Nombre</label> <div class=\"control\"><input type=\"text\" placeholder=\"Text input\"" + _vm._ssrAttr("value", _vm.dogName) + " class=\"input\"></div></div> <div class=\"field\"><label class=\"label\">Raza</label> <div class=\"control\"><input type=\"text\" placeholder=\"Text input\"" + _vm._ssrAttr("value", _vm.dogBreed) + " class=\"input\"></div></div> <div class=\"field\"><label class=\"label\">Color</label> <div class=\"control\"><input type=\"text\" placeholder=\"Text input\"" + _vm._ssrAttr("value", _vm.dogColor) + " class=\"input\"></div></div> <div class=\"field\"><label class=\"label\">Edad</label> <div class=\"control\"><input type=\"text\" placeholder=\"Text input\"" + _vm._ssrAttr("value", _vm.dogAge) + " class=\"input\"></div></div> "), _vm._ssrNode("<div class=\"field\">", "</div>", [_vm._ssrNode("<label class=\"label\">Selecciona un dueño:</label> "), _c('select', {
    directives: [{
      name: "model",
      rawName: "v-model",
      value: _vm.selectedOwner,
      expression: "selectedOwner"
    }],
    on: {
      "change": function ($event) {
        var $$selectedVal = Array.prototype.filter.call($event.target.options, function (o) {
          return o.selected;
        }).map(function (o) {
          var val = "_value" in o ? o._value : o.value;
          return val;
        });
        _vm.selectedOwner = $event.target.multiple ? $$selectedVal : $$selectedVal[0];
      }
    }
  }, _vm._l(_vm.owners, function (Owner) {
    return _c('option', {
      key: Owner.id,
      domProps: {
        "value": Owner.id
      }
    }, [_vm._v("\n                    " + _vm._s(Owner.Nombre) + "\n                ")]);
  }), 0)], 2), _vm._ssrNode(" <div class=\"field is-grouped\"><div class=\"control\"><button id=\"button\" class=\"button is-link\">Submit</button></div> <div class=\"control\"><button id=\"buttonInv\" class=\"button is-link is-light\">Cancel</button></div></div>")], 2);
};
var staticRenderFns = [];

// CONCATENATED MODULE: ./components/DogRegister.vue?vue&type=template&id=4a236bd2&

// CONCATENATED MODULE: ./node_modules/babel-loader/lib??ref--2-0!./node_modules/@nuxt/components/dist/loader.js??ref--0-0!./node_modules/vue-loader/lib??vue-loader-options!./components/DogRegister.vue?vue&type=script&lang=js&
/* harmony default export */ var DogRegistervue_type_script_lang_js_ = ({
  data() {
    return {
      dogName: '',
      dogAge: '',
      dogColor: '',
      dogBreed: '',
      owners: [],
      selectedOwner: null
    };
  },
  methods: {
    async postDogData() {
      const dogData = {
        Nombre: this.dogName,
        Edad: this.dogAge,
        Color: this.dogColor,
        Raza: this.dogBreed,
        Dueño_id: parseInt(this.selectedOwner)
      };
      const path = "http://localhost:8000" + '/perro/';
      try {
        const res = await this.$axios.post(path, dogData);
      } catch (error) {
        console.log(error);
      }
    },
    async getOwners() {
      try {
        const path = "http://localhost:8000" + '/dueno/';
        const response = await this.$axios.get(path);
        this.owners = response.data;
      } catch (error) {
        console.error('Error al obtener la lista de dueños:', error);
      }
    },
    clear() {
      this.dogName = '', this.dogAge = '', this.dogColor = '', this.dogBreed = '', this.selectedOwner = null;
    }
  },
  async created() {
    await this.getOwners();
  }
});
// CONCATENATED MODULE: ./components/DogRegister.vue?vue&type=script&lang=js&
 /* harmony default export */ var components_DogRegistervue_type_script_lang_js_ = (DogRegistervue_type_script_lang_js_); 
// EXTERNAL MODULE: ./node_modules/vue-loader/lib/runtime/componentNormalizer.js
var componentNormalizer = __webpack_require__(2);

// CONCATENATED MODULE: ./components/DogRegister.vue



function injectStyles (context) {
  
  var style0 = __webpack_require__(39)
if (style0.__inject__) style0.__inject__(context)

}

/* normalize component */

var component = Object(componentNormalizer["a" /* default */])(
  components_DogRegistervue_type_script_lang_js_,
  render,
  staticRenderFns,
  false,
  injectStyles,
  null,
  "340938bf"
  
)

/* harmony default export */ var DogRegister = __webpack_exports__["default"] = (component.exports);

/***/ }),

/***/ 54:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
// ESM COMPAT FLAG
__webpack_require__.r(__webpack_exports__);

// CONCATENATED MODULE: ./node_modules/babel-loader/lib??ref--2-0!./node_modules/vue-loader/lib/loaders/templateLoader.js??ref--6!./node_modules/@nuxt/components/dist/loader.js??ref--0-0!./node_modules/vue-loader/lib??vue-loader-options!./pages/nuevos.vue?vue&type=template&id=01aa98a3&
var render = function render() {
  var _vm = this,
    _c = _vm._self._c;
  return _c('div', {
    staticClass: "container"
  }, [_c('DogRegister')], 1);
};
var staticRenderFns = [];

// CONCATENATED MODULE: ./pages/nuevos.vue?vue&type=template&id=01aa98a3&

// EXTERNAL MODULE: ./components/DogRegister.vue + 4 modules
var DogRegister = __webpack_require__(45);

// CONCATENATED MODULE: ./node_modules/babel-loader/lib??ref--2-0!./node_modules/@nuxt/components/dist/loader.js??ref--0-0!./node_modules/vue-loader/lib??vue-loader-options!./pages/nuevos.vue?vue&type=script&lang=js&

/* harmony default export */ var nuevosvue_type_script_lang_js_ = ({
  data() {
    return {};
  },
  components: {
    DogRegister: DogRegister["default"]
  }
});
// CONCATENATED MODULE: ./pages/nuevos.vue?vue&type=script&lang=js&
 /* harmony default export */ var pages_nuevosvue_type_script_lang_js_ = (nuevosvue_type_script_lang_js_); 
// EXTERNAL MODULE: ./node_modules/vue-loader/lib/runtime/componentNormalizer.js
var componentNormalizer = __webpack_require__(2);

// CONCATENATED MODULE: ./pages/nuevos.vue





/* normalize component */

var component = Object(componentNormalizer["a" /* default */])(
  pages_nuevosvue_type_script_lang_js_,
  render,
  staticRenderFns,
  false,
  null,
  null,
  "5954774c"
  
)

/* harmony default export */ var nuevos = __webpack_exports__["default"] = (component.exports);

/* nuxt-component-imports */
installComponents(component, {DogRegister: __webpack_require__(45).default})


/***/ })

};;
//# sourceMappingURL=nuevos.js.map