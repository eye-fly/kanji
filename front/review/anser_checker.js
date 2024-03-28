!function(t) {
    var e = {};
    function n(r) {
        if (e[r])
            return e[r].exports;
        var o = e[r] = {
            i: r,
            l: !1,
            exports: {}
        };
        return t[r].call(o.exports, o, o.exports, n), o.l = !0, o.exports
    }
    n.m = t,
    n.c = e,
    n.d = function(t, e, r) {
        n.o(t, e) || Object.defineProperty(t, e, {
            enumerable: !0,
            get: r
        })
    },
    n.r = function(t) {
        "undefined" !== typeof Symbol && Symbol.toStringTag && Object.defineProperty(t, Symbol.toStringTag, {
            value: "Module"
        }),
        Object.defineProperty(t, "__esModule", {
            value: !0
        })
    },
    n.t = function(t, e) {
        if (1 & e && (t = n(t)), 8 & e)
            return t;
        if (4 & e && "object" === typeof t && t && t.__esModule)
            return t;
        var r = Object.create(null);
        if (n.r(r), Object.defineProperty(r, "default", {
            enumerable: !0,
            value: t
        }), 2 & e && "string" != typeof t)
            for (var o in t)
                n.d(r, o, function(e) {
                    return t[e]
                }.bind(null, o));
        return r
    },
    n.n = function(t) {
        var e = t && t.__esModule ? function() {
            return t.default
        } : function() {
            return t
        };
        return n.d(e, "a", e), e
    },
    n.o = function(t, e) {
        return Object.prototype.hasOwnProperty.call(t, e)
    },
    n.p = "https://assets.wanikani.com/packs/",
    n(n.s = 976)
}({
    161: function(t, e, n) {
        "use strict";
        var r = {
            onyomi: "on\u2019yomi",
            kunyomi: "kun\u2019yomi",
            nanori: "nanori"
        };
        e.a = function(t) {
            return r[t]
        }
    },
    26: function(t, e, n) {
        (function(n) {
            var r,
                o,
                i,
                u;
            function a(t) {
                return (a = "function" === typeof Symbol && "symbol" === typeof Symbol.iterator ? function(t) {
                    return typeof t
                } : function(t) {
                    return t && "function" === typeof Symbol && t.constructor === Symbol && t !== Symbol.prototype ? "symbol" : typeof t
                })(t)
            }
            u = function(t) {
                "use strict";
                var e = "undefined" != typeof window ? window : "undefined" != typeof n ? n : "undefined" != typeof self ? self : {};
                function r(t, e) {
                    return t(e = {
                        exports: {}
                    }, e.exports), e.exports
                }
                var o = r((function(t) {
                        var e = t.exports = "undefined" != typeof window && window.Math == Math ? window : "undefined" != typeof self && self.Math == Math ? self : Function("return this")();
                        "number" == typeof __g && (__g = e)
                    })),
                    i = Object.freeze({
                        default: o,
                        __moduleExports: o
                    }),
                    u = r((function(t) {
                        var e = t.exports = {
                            version: "2.5.5"
                        };
                        "number" == typeof __e && (__e = e)
                    })),
                    c = Object.freeze({
                        default: u,
                        __moduleExports: u,
                        version: u.version
                    }),
                    f = function(t) {
                        return "object" == a(t) ? null !== t : "function" == typeof t
                    },
                    s = Object.freeze({
                        default: f,
                        __moduleExports: f
                    }),
                    l = s && f || s,
                    h = function(t) {
                        if (!l(t))
                            throw TypeError(t + " is not an object!");
                        return t
                    },
                    d = Object.freeze({
                        default: h,
                        __moduleExports: h
                    }),
                    v = function(t) {
                        try {
                            return !!t()
                        } catch (t) {
                            return !0
                        }
                    },
                    p = Object.freeze({
                        default: v,
                        __moduleExports: v
                    }),
                    y = p && v || p,
                    g = !y((function() {
                        return 7 != Object.defineProperty({}, "a", {
                            get: function() {
                                return 7
                            }
                        }).a
                    })),
                    m = Object.freeze({
                        default: g,
                        __moduleExports: g
                    }),
                    b = i && o || i,
                    _ = b.document,
                    O = l(_) && l(_.createElement),
                    j = function(t) {
                        return O ? _.createElement(t) : {}
                    },
                    E = Object.freeze({
                        default: j,
                        __moduleExports: j
                    }),
                    w = m && g || m,
                    x = E && j || E,
                    S = !w && !y((function() {
                        return 7 != Object.defineProperty(x("div"), "a", {
                            get: function() {
                                return 7
                            }
                        }).a
                    })),
                    M = Object.freeze({
                        default: S,
                        __moduleExports: S
                    }),
                    z = function(t, e) {
                        if (!l(t))
                            return t;
                        var n,
                            r;
                        if (e && "function" == typeof (n = t.toString) && !l(r = n.call(t)))
                            return r;
                        if ("function" == typeof (n = t.valueOf) && !l(r = n.call(t)))
                            return r;
                        if (!e && "function" == typeof (n = t.toString) && !l(r = n.call(t)))
                            return r;
                        throw TypeError("Can't convert object to primitive value")
                    },
                    A = Object.freeze({
                        default: z,
                        __moduleExports: z
                    }),
                    P = d && h || d,
                    k = M && S || M,
                    L = A && z || A,
                    F = Object.defineProperty,
                    N = w ? Object.defineProperty : function(t, e, n) {
                        if (P(t), e = L(e, !0), P(n), k)
                            try {
                                return F(t, e, n)
                            } catch (t) {}
                        if ("get" in n || "set" in n)
                            throw TypeError("Accessors not supported!");
                        return "value" in n && (t[e] = n.value), t
                    },
                    I = {
                        f: N
                    },
                    T = Object.freeze({
                        default: I,
                        __moduleExports: I,
                        f: N
                    }),
                    R = function(t, e) {
                        return {
                            enumerable: !(1 & t),
                            configurable: !(2 & t),
                            writable: !(4 & t),
                            value: e
                        }
                    },
                    C = Object.freeze({
                        default: R,
                        __moduleExports: R
                    }),
                    W = T && I || T,
                    D = C && R || C,
                    B = w ? function(t, e, n) {
                        return W.f(t, e, D(1, n))
                    } : function(t, e, n) {
                        return t[e] = n, t
                    },
                    K = Object.freeze({
                        default: B,
                        __moduleExports: B
                    }),
                    V = {}.hasOwnProperty,
                    U = function(t, e) {
                        return V.call(t, e)
                    },
                    G = Object.freeze({
                        default: U,
                        __moduleExports: U
                    }),
                    H = 0,
                    J = Math.random(),
                    Y = function(t) {
                        return "Symbol(".concat(void 0 === t ? "" : t, ")_", (++H + J).toString(36))
                    },
                    q = Object.freeze({
                        default: Y,
                        __moduleExports: Y
                    }),
                    X = K && B || K,
                    $ = G && U || G,
                    Q = q && Y || q,
                    Z = c && u || c,
                    tt = r((function(t) {
                        var e = Q("src"),
                            n = Function.toString,
                            r = ("" + n).split("toString");
                        Z.inspectSource = function(t) {
                            return n.call(t)
                        },
                        (t.exports = function(t, n, o, i) {
                            var u = "function" == typeof o;
                            u && ($(o, "name") || X(o, "name", n)),
                            t[n] !== o && (u && ($(o, e) || X(o, e, t[n] ? "" + t[n] : r.join(n + ""))), t === b ? t[n] = o : i ? t[n] ? t[n] = o : X(t, n, o) : (delete t[n], X(t, n, o)))
                        })(Function.prototype, "toString", (function() {
                            return "function" == typeof this && this[e] || n.call(this)
                        }))
                    })),
                    et = Object.freeze({
                        default: tt,
                        __moduleExports: tt
                    }),
                    nt = function(t) {
                        if ("function" != typeof t)
                            throw TypeError(t + " is not a function!");
                        return t
                    },
                    rt = Object.freeze({
                        default: nt,
                        __moduleExports: nt
                    }),
                    ot = rt && nt || rt,
                    it = function(t, e, n) {
                        if (ot(t), void 0 === e)
                            return t;
                        switch (n) {
                        case 1:
                            return function(n) {
                                return t.call(e, n)
                            };
                        case 2:
                            return function(n, r) {
                                return t.call(e, n, r)
                            };
                        case 3:
                            return function(n, r, o) {
                                return t.call(e, n, r, o)
                            }
                        }
                        return function() {
                            return t.apply(e, arguments)
                        }
                    },
                    ut = Object.freeze({
                        default: it,
                        __moduleExports: it
                    }),
                    at = et && tt || et,
                    ct = ut && it || ut,
                    ft = function t(e, n, r) {
                        var o,
                            i,
                            u,
                            a,
                            c = e & t.F,
                            f = e & t.G,
                            s = e & t.P,
                            l = e & t.B,
                            h = f ? b : e & t.S ? b[n] || (b[n] = {}) : (b[n] || {}).prototype,
                            d = f ? Z : Z[n] || (Z[n] = {}),
                            v = d.prototype || (d.prototype = {});
                        for (o in f && (r = n), r)
                            u = ((i = !c && h && void 0 !== h[o]) ? h : r)[o],
                            a = l && i ? ct(u, b) : s && "function" == typeof u ? ct(Function.call, u) : u,
                            h && at(h, o, u, e & t.U),
                            d[o] != u && X(d, o, a),
                            s && v[o] != u && (v[o] = u)
                    };
                b.core = Z,
                ft.F = 1,
                ft.G = 2,
                ft.S = 4,
                ft.P = 8,
                ft.B = 16,
                ft.W = 32,
                ft.U = 64,
                ft.R = 128;
                for (var st, lt = ft, ht = Object.freeze({
                        default: lt,
                        __moduleExports: lt
                    }), dt = Q("typed_array"), vt = Q("view"), pt = !(!b.ArrayBuffer || !b.DataView), yt = pt, gt = 0, mt = "Int8Array,Uint8Array,Uint8ClampedArray,Int16Array,Uint16Array,Int32Array,Uint32Array,Float32Array,Float64Array".split(","); 9 > gt;)
                    (st = b[mt[gt++]]) ? (X(st.prototype, dt, !0), X(st.prototype, vt, !0)) : yt = !1;
                var bt = {
                        ABV: pt,
                        CONSTR: yt,
                        TYPED: dt,
                        VIEW: vt
                    },
                    _t = Object.freeze({
                        default: bt,
                        __moduleExports: bt,
                        ABV: bt.ABV,
                        CONSTR: bt.CONSTR,
                        TYPED: bt.TYPED,
                        VIEW: bt.VIEW
                    }),
                    Ot = Object.freeze({
                        default: !1,
                        __moduleExports: !1
                    }),
                    jt = function(t, e, n) {
                        for (var r in e)
                            at(t, r, e[r], n);
                        return t
                    },
                    Et = Object.freeze({
                        default: jt,
                        __moduleExports: jt
                    }),
                    wt = function(t, e, n, r) {
                        if (!(t instanceof e) || void 0 !== r && r in t)
                            throw TypeError(n + ": incorrect invocation!");
                        return t
                    },
                    xt = Object.freeze({
                        default: wt,
                        __moduleExports: wt
                    }),
                    St = Math.ceil,
                    Mt = Math.floor,
                    zt = function(t) {
                        return isNaN(t = +t) ? 0 : (t > 0 ? Mt : St)(t)
                    },
                    At = Object.freeze({
                        default: zt,
                        __moduleExports: zt
                    }),
                    Pt = At && zt || At,
                    kt = Math.min,
                    Lt = function(t) {
                        return t > 0 ? kt(Pt(t), 9007199254740991) : 0
                    },
                    Ft = Object.freeze({
                        default: Lt,
                        __moduleExports: Lt
                    }),
                    Nt = Ft && Lt || Ft,
                    It = function(t) {
                        if (void 0 === t)
                            return 0;
                        var e = Pt(t),
                            n = Nt(e);
                        if (e !== n)
                            throw RangeError("Wrong length!");
                        return n
                    },
                    Tt = Object.freeze({
                        default: It,
                        __moduleExports: It
                    }),
                    Rt = {}.toString,
                    Ct = function(t) {
                        return Rt.call(t).slice(8, -1)
                    },
                    Wt = Object.freeze({
                        default: Ct,
                        __moduleExports: Ct
                    }),
                    Dt = Wt && Ct || Wt,
                    Bt = Object("z").propertyIsEnumerable(0) ? Object : function(t) {
                        return "String" == Dt(t) ? t.split("") : Object(t)
                    },
                    Kt = Object.freeze({
                        default: Bt,
                        __moduleExports: Bt
                    }),
                    Vt = function(t) {
                        if (void 0 == t)
                            throw TypeError("Can't call method on  " + t);
                        return t
                    },
                    Ut = Object.freeze({
                        default: Vt,
                        __moduleExports: Vt
                    }),
                    Gt = Kt && Bt || Kt,
                    Ht = Ut && Vt || Ut,
                    Jt = function(t) {
                        return Gt(Ht(t))
                    },
                    Yt = Object.freeze({
                        default: Jt,
                        __moduleExports: Jt
                    }),
                    qt = Math.max,
                    Xt = Math.min,
                    $t = function(t, e) {
                        return 0 > (t = Pt(t)) ? qt(t + e, 0) : Xt(t, e)
                    },
                    Qt = Object.freeze({
                        default: $t,
                        __moduleExports: $t
                    }),
                    Zt = Yt && Jt || Yt,
                    te = Qt && $t || Qt,
                    ee = function(t) {
                        return function(e, n, r) {
                            var o,
                                i = Zt(e),
                                u = Nt(i.length),
                                a = te(r, u);
                            if (t && n != n) {
                                for (; u > a;)
                                    if ((o = i[a++]) != o)
                                        return !0
                            } else
                                for (; u > a; a++)
                                    if ((t || a in i) && i[a] === n)
                                        return t || a || 0;
                            return !t && -1
                        }
                    },
                    ne = Object.freeze({
                        default: ee,
                        __moduleExports: ee
                    }),
                    re = b["__core-js_shared__"] || (b["__core-js_shared__"] = {}),
                    oe = function(t) {
                        return re[t] || (re[t] = {})
                    },
                    ie = Object.freeze({
                        default: oe,
                        __moduleExports: oe
                    }),
                    ue = ie && oe || ie,
                    ae = ue("keys"),
                    ce = function(t) {
                        return ae[t] || (ae[t] = Q(t))
                    },
                    fe = Object.freeze({
                        default: ce,
                        __moduleExports: ce
                    }),
                    se = ne && ee || ne,
                    le = fe && ce || fe,
                    he = se(!1),
                    de = le("IE_PROTO"),
                    ve = function(t, e) {
                        var n,
                            r = Zt(t),
                            o = 0,
                            i = [];
                        for (n in r)
                            n != de && $(r, n) && i.push(n);
                        for (; e.length > o;)
                            $(r, n = e[o++]) && (~he(i, n) || i.push(n));
                        return i
                    },
                    pe = Object.freeze({
                        default: ve,
                        __moduleExports: ve
                    }),
                    ye = "constructor,hasOwnProperty,isPrototypeOf,propertyIsEnumerable,toLocaleString,toString,valueOf".split(","),
                    ge = Object.freeze({
                        default: ye,
                        __moduleExports: ye
                    }),
                    me = pe && ve || pe,
                    be = ge && ye || ge,
                    _e = be.concat("length", "prototype"),
                    Oe = Object.getOwnPropertyNames || function(t) {
                        return me(t, _e)
                    },
                    je = {
                        f: Oe
                    },
                    Ee = Object.freeze({
                        default: je,
                        __moduleExports: je,
                        f: Oe
                    }),
                    we = function(t) {
                        return Object(Ht(t))
                    },
                    xe = Object.freeze({
                        default: we,
                        __moduleExports: we
                    }),
                    Se = xe && we || xe,
                    Me = function(t) {
                        for (var e = Se(this), n = Nt(e.length), r = arguments.length, o = te(r > 1 ? arguments[1] : void 0, n), i = r > 2 ? arguments[2] : void 0, u = void 0 === i ? n : te(i, n); u > o;)
                            e[o++] = t;
                        return e
                    },
                    ze = Object.freeze({
                        default: Me,
                        __moduleExports: Me
                    }),
                    Ae = r((function(t) {
                        var e = ue("wks"),
                            n = b.Symbol,
                            r = "function" == typeof n;
                        (t.exports = function(t) {
                            return e[t] || (e[t] = r && n[t] || (r ? n : Q)("Symbol." + t))
                        }).store = e
                    })),
                    Pe = Object.freeze({
                        default: Ae,
                        __moduleExports: Ae
                    }),
                    ke = Pe && Ae || Pe,
                    Le = W.f,
                    Fe = ke("toStringTag"),
                    Ne = function(t, e, n) {
                        t && !$(t = n ? t : t.prototype, Fe) && Le(t, Fe, {
                            configurable: !0,
                            value: e
                        })
                    },
                    Ie = Object.freeze({
                        default: Ne,
                        __moduleExports: Ne
                    }),
                    Te = Ot,
                    Re = _t && bt || _t,
                    Ce = Et && jt || Et,
                    We = xt && wt || xt,
                    De = Tt && It || Tt,
                    Be = Ee && je || Ee,
                    Ke = ze && Me || ze,
                    Ve = Ie && Ne || Ie,
                    Ue = r((function(t, e) {
                        var n = Be.f,
                            r = W.f,
                            o = "prototype",
                            i = "Wrong index!",
                            u = b.ArrayBuffer,
                            a = b.DataView,
                            c = b.Math,
                            f = b.RangeError,
                            s = b.Infinity,
                            l = u,
                            h = c.abs,
                            d = c.pow,
                            v = c.floor,
                            p = c.log,
                            g = c.LN2,
                            m = w ? "_b" : "buffer",
                            _ = w ? "_l" : "byteLength",
                            O = w ? "_o" : "byteOffset";
                        function j(t, e, n) {
                            var r,
                                o,
                                i,
                                u = Array(n),
                                a = 8 * n - e - 1,
                                c = (1 << a) - 1,
                                f = c >> 1,
                                l = 23 === e ? d(2, -24) - d(2, -77) : 0,
                                y = 0,
                                m = 0 > t || 0 === t && 0 > 1 / t ? 1 : 0;
                            for ((t = h(t)) != t || t === s ? (o = t != t ? 1 : 0, r = c) : (r = v(p(t) / g), 1 > t * (i = d(2, -r)) && (r--, i *= 2), 2 > (t += 1 > r + f ? l * d(2, 1 - f) : l / i) * i || (r++, i /= 2), c > r + f ? 1 > r + f ? (o = t * d(2, f - 1) * d(2, e), r = 0) : (o = (t * i - 1) * d(2, e), r += f) : (o = 0, r = c)); e >= 8; u[y++] = 255 & o, o /= 256, e -= 8)
                                ;
                            for (r = r << e | o, a += e; a > 0; u[y++] = 255 & r, r /= 256, a -= 8)
                                ;
                            return u[--y] |= 128 * m, u
                        }
                        function E(t, e, n) {
                            var r,
                                o = 8 * n - e - 1,
                                i = (1 << o) - 1,
                                u = i >> 1,
                                a = o - 7,
                                c = n - 1,
                                f = t[c--],
                                l = 127 & f;
                            for (f >>= 7; a > 0; l = 256 * l + t[c], c--, a -= 8)
                                ;
                            for (r = l & (1 << -a) - 1, l >>= -a, a += e; a > 0; r = 256 * r + t[c], c--, a -= 8)
                                ;
                            if (0 === l)
                                l = 1 - u;
                            else {
                                if (l === i)
                                    return r ? NaN : f ? -s : s;
                                r += d(2, e),
                                l -= u
                            }
                            return (f ? -1 : 1) * r * d(2, l - e)
                        }
                        function x(t) {
                            return t[3] << 24 | t[2] << 16 | t[1] << 8 | t[0]
                        }
                        function S(t) {
                            return [255 & t]
                        }
                        function M(t) {
                            return [255 & t, t >> 8 & 255]
                        }
                        function z(t) {
                            return [255 & t, t >> 8 & 255, t >> 16 & 255, t >> 24 & 255]
                        }
                        function A(t) {
                            return j(t, 52, 8)
                        }
                        function P(t) {
                            return j(t, 23, 4)
                        }
                        function k(t, e, n) {
                            r(t[o], e, {
                                get: function() {
                                    return this[n]
                                }
                            })
                        }
                        function L(t, e, n, r) {
                            var o = De(+n);
                            if (o + e > t[_])
                                throw f(i);
                            var u = o + t[O],
                                a = t[m]._b.slice(u, u + e);
                            return r ? a : a.reverse()
                        }
                        function F(t, e, n, r, o, u) {
                            var a = De(+n);
                            if (a + e > t[_])
                                throw f(i);
                            for (var c = t[m]._b, s = a + t[O], l = r(+o), h = 0; e > h; h++)
                                c[s + h] = l[u ? h : e - h - 1]
                        }
                        if (Re.ABV) {
                            if (!y((function() {
                                u(1)
                            })) || !y((function() {
                                new u(-1)
                            })) || y((function() {
                                return new u, new u(1.5), new u(NaN), "ArrayBuffer" != u.name
                            }))) {
                                for (var N, I = (u = function(t) {
                                        return We(this, u), new l(De(t))
                                    })[o] = l[o], T = n(l), R = 0; T.length > R;)
                                    (N = T[R++]) in u || X(u, N, l[N]);
                                Te || (I.constructor = u)
                            }
                            var C = new a(new u(2)),
                                D = a[o].setInt8;
                            C.setInt8(0, 2147483648),
                            C.setInt8(1, 2147483649),
                            !C.getInt8(0) && C.getInt8(1) || Ce(a[o], {
                                setInt8: function(t, e) {
                                    D.call(this, t, e << 24 >> 24)
                                },
                                setUint8: function(t, e) {
                                    D.call(this, t, e << 24 >> 24)
                                }
                            }, !0)
                        } else
                            u = function(t) {
                                We(this, u, "ArrayBuffer");
                                var e = De(t);
                                this._b = Ke.call(Array(e), 0),
                                this[_] = e
                            },
                            a = function(t, e, n) {
                                We(this, a, "DataView"),
                                We(t, u, "DataView");
                                var r = t[_],
                                    o = Pt(e);
                                if (0 > o || o > r)
                                    throw f("Wrong offset!");
                                if (o + (n = void 0 === n ? r - o : Nt(n)) > r)
                                    throw f("Wrong length!");
                                this[m] = t,
                                this[O] = o,
                                this[_] = n
                            },
                            w && (k(u, "byteLength", "_l"), k(a, "buffer", "_b"), k(a, "byteLength", "_l"), k(a, "byteOffset", "_o")),
                            Ce(a[o], {
                                getInt8: function(t) {
                                    return L(this, 1, t)[0] << 24 >> 24
                                },
                                getUint8: function(t) {
                                    return L(this, 1, t)[0]
                                },
                                getInt16: function(t) {
                                    var e = L(this, 2, t, arguments[1]);
                                    return (e[1] << 8 | e[0]) << 16 >> 16
                                },
                                getUint16: function(t) {
                                    var e = L(this, 2, t, arguments[1]);
                                    return e[1] << 8 | e[0]
                                },
                                getInt32: function(t) {
                                    return x(L(this, 4, t, arguments[1]))
                                },
                                getUint32: function(t) {
                                    return x(L(this, 4, t, arguments[1])) >>> 0
                                },
                                getFloat32: function(t) {
                                    return E(L(this, 4, t, arguments[1]), 23, 4)
                                },
                                getFloat64: function(t) {
                                    return E(L(this, 8, t, arguments[1]), 52, 8)
                                },
                                setInt8: function(t, e) {
                                    F(this, 1, t, S, e)
                                },
                                setUint8: function(t, e) {
                                    F(this, 1, t, S, e)
                                },
                                setInt16: function(t, e) {
                                    F(this, 2, t, M, e, arguments[2])
                                },
                                setUint16: function(t, e) {
                                    F(this, 2, t, M, e, arguments[2])
                                },
                                setInt32: function(t, e) {
                                    F(this, 4, t, z, e, arguments[2])
                                },
                                setUint32: function(t, e) {
                                    F(this, 4, t, z, e, arguments[2])
                                },
                                setFloat32: function(t, e) {
                                    F(this, 4, t, P, e, arguments[2])
                                },
                                setFloat64: function(t, e) {
                                    F(this, 8, t, A, e, arguments[2])
                                }
                            });
                        Ve(u, "ArrayBuffer"),
                        Ve(a, "DataView"),
                        X(a[o], Re.VIEW, !0),
                        e.ArrayBuffer = u,
                        e.DataView = a
                    })),
                    Ge = Object.freeze({
                        default: Ue,
                        __moduleExports: Ue
                    }),
                    He = ke("species"),
                    Je = function(t, e) {
                        var n,
                            r = P(t).constructor;
                        return void 0 === r || void 0 == (n = P(r)[He]) ? e : ot(n)
                    },
                    Ye = Object.freeze({
                        default: Je,
                        __moduleExports: Je
                    }),
                    qe = ke("species"),
                    Xe = function(t) {
                        var e = b[t];
                        w && e && !e[qe] && W.f(e, qe, {
                            configurable: !0,
                            get: function() {
                                return this
                            }
                        })
                    },
                    $e = Object.freeze({
                        default: Xe,
                        __moduleExports: Xe
                    }),
                    Qe = ht && lt || ht,
                    Ze = Ge && Ue || Ge,
                    tn = Ye && Je || Ye,
                    en = $e && Xe || $e,
                    nn = b.ArrayBuffer,
                    rn = Ze.ArrayBuffer,
                    on = Ze.DataView,
                    un = Re.ABV && nn.isView,
                    an = rn.prototype.slice,
                    cn = Re.VIEW;
                Qe(Qe.G + Qe.W + Qe.F * (nn !== rn), {
                    ArrayBuffer: rn
                }),
                Qe(Qe.S + Qe.F * !Re.CONSTR, "ArrayBuffer", {
                    isView: function(t) {
                        return un && un(t) || l(t) && cn in t
                    }
                }),
                Qe(Qe.P + Qe.U + Qe.F * y((function() {
                    return !new rn(2).slice(1, void 0).byteLength
                })), "ArrayBuffer", {
                    slice: function(t, e) {
                        if (void 0 !== an && void 0 === e)
                            return an.call(P(this), t);
                        for (var n = P(this).byteLength, r = te(t, n), o = te(void 0 === e ? n : e, n), i = new (tn(this, rn))(Nt(o - r)), u = new on(this), a = new on(i), c = 0; o > r;)
                            a.setUint8(c++, u.getUint8(r++));
                        return i
                    }
                }),
                en("ArrayBuffer");
                var fn = ke("toStringTag"),
                    sn = "Arguments" == Dt(function() {
                        return arguments
                    }()),
                    ln = function(t) {
                        var e,
                            n,
                            r;
                        return void 0 === t ? "Undefined" : null === t ? "Null" : "string" == typeof (n = function(t, e) {
                            try {
                                return t[e]
                            } catch (t) {}
                        }(e = Object(t), fn)) ? n : sn ? Dt(e) : "Object" == (r = Dt(e)) && "function" == typeof e.callee ? "Arguments" : r
                    },
                    hn = Object.freeze({
                        default: ln,
                        __moduleExports: ln
                    }),
                    dn = {},
                    vn = Object.freeze({
                        default: dn,
                        __moduleExports: dn
                    }),
                    pn = vn && dn || vn,
                    yn = ke("iterator"),
                    gn = Array.prototype,
                    mn = function(t) {
                        return void 0 !== t && (pn.Array === t || gn[yn] === t)
                    },
                    bn = Object.freeze({
                        default: mn,
                        __moduleExports: mn
                    }),
                    _n = Object.keys || function(t) {
                        return me(t, be)
                    },
                    On = Object.freeze({
                        default: _n,
                        __moduleExports: _n
                    }),
                    jn = On && _n || On,
                    En = w ? Object.defineProperties : function(t, e) {
                        P(t);
                        for (var n, r = jn(e), o = r.length, i = 0; o > i;)
                            W.f(t, n = r[i++], e[n]);
                        return t
                    },
                    wn = Object.freeze({
                        default: En,
                        __moduleExports: En
                    }),
                    xn = b.document,
                    Sn = xn && xn.documentElement,
                    Mn = Object.freeze({
                        default: Sn,
                        __moduleExports: Sn
                    }),
                    zn = wn && En || wn,
                    An = Mn && Sn || Mn,
                    Pn = le("IE_PROTO"),
                    kn = function() {},
                    Ln = function() {
                        var t,
                            e = x("iframe"),
                            n = be.length;
                        for (e.style.display = "none", An.appendChild(e), e.src = "javascript:", (t = e.contentWindow.document).open(), t.write("<script>document.F=Object<\/script>"), t.close(), Ln = t.F; n--;)
                            delete Ln.prototype[be[n]];
                        return Ln()
                    },
                    Fn = Object.create || function(t, e) {
                        var n;
                        return null !== t ? (kn.prototype = P(t), n = new kn, kn.prototype = null, n[Pn] = t) : n = Ln(), void 0 === e ? n : zn(n, e)
                    },
                    Nn = Object.freeze({
                        default: Fn,
                        __moduleExports: Fn
                    }),
                    In = le("IE_PROTO"),
                    Tn = Object.prototype,
                    Rn = Object.getPrototypeOf || function(t) {
                        return t = Se(t), $(t, In) ? t[In] : "function" == typeof t.constructor && t instanceof t.constructor ? t.constructor.prototype : t instanceof Object ? Tn : null
                    },
                    Cn = Object.freeze({
                        default: Rn,
                        __moduleExports: Rn
                    }),
                    Wn = hn && ln || hn,
                    Dn = ke("iterator"),
                    Bn = Z.getIteratorMethod = function(t) {
                        if (void 0 != t)
                            return t[Dn] || t["@@iterator"] || pn[Wn(t)]
                    },
                    Kn = Object.freeze({
                        default: Bn,
                        __moduleExports: Bn
                    }),
                    Vn = Array.isArray || function(t) {
                        return "Array" == Dt(t)
                    },
                    Un = Object.freeze({
                        default: Vn,
                        __moduleExports: Vn
                    }),
                    Gn = Un && Vn || Un,
                    Hn = ke("species"),
                    Jn = function(t) {
                        var e;
                        return Gn(t) && ("function" != typeof (e = t.constructor) || e !== Array && !Gn(e.prototype) || (e = void 0), l(e) && null === (e = e[Hn]) && (e = void 0)), void 0 === e ? Array : e
                    },
                    Yn = Object.freeze({
                        default: Jn,
                        __moduleExports: Jn
                    }),
                    qn = Yn && Jn || Yn,
                    Xn = function(t, e) {
                        return new (qn(t))(e)
                    },
                    $n = Object.freeze({
                        default: Xn,
                        __moduleExports: Xn
                    }),
                    Qn = $n && Xn || $n,
                    Zn = function(t, e) {
                        var n = 1 == t,
                            r = 2 == t,
                            o = 3 == t,
                            i = 4 == t,
                            u = 6 == t,
                            a = 5 == t || u,
                            c = e || Qn;
                        return function(e, f, s) {
                            for (var l, h, d = Se(e), v = Gt(d), p = ct(f, s, 3), y = Nt(v.length), g = 0, m = n ? c(e, y) : r ? c(e, 0) : void 0; y > g; g++)
                                if ((a || g in v) && (h = p(l = v[g], g, d), t))
                                    if (n)
                                        m[g] = h;
                                    else if (h)
                                        switch (t) {
                                        case 3:
                                            return !0;
                                        case 5:
                                            return l;
                                        case 6:
                                            return g;
                                        case 2:
                                            m.push(l)
                                        }
                                    else if (i)
                                        return !1;
                            return u ? -1 : o || i ? i : m
                        }
                    },
                    tr = Object.freeze({
                        default: Zn,
                        __moduleExports: Zn
                    }),
                    er = ke("unscopables"),
                    nr = Array.prototype;
                void 0 == nr[er] && X(nr, er, {});
                var rr = function(t) {
                        nr[er][t] = !0
                    },
                    or = Object.freeze({
                        default: rr,
                        __moduleExports: rr
                    }),
                    ir = function(t, e) {
                        return {
                            value: e,
                            done: !!t
                        }
                    },
                    ur = Object.freeze({
                        default: ir,
                        __moduleExports: ir
                    }),
                    ar = Nn && Fn || Nn,
                    cr = {};
                X(cr, ke("iterator"), (function() {
                    return this
                }));
                var fr = function(t, e, n) {
                        t.prototype = ar(cr, {
                            next: D(1, n)
                        }),
                        Ve(t, e + " Iterator")
                    },
                    sr = Object.freeze({
                        default: fr,
                        __moduleExports: fr
                    }),
                    lr = sr && fr || sr,
                    hr = Cn && Rn || Cn,
                    dr = ke("iterator"),
                    vr = !([].keys && "next" in [].keys()),
                    pr = function() {
                        return this
                    },
                    yr = function(t, e, n, r, o, i, u) {
                        lr(n, e, r);
                        var a,
                            c,
                            f,
                            s = function(t) {
                                if (!vr && t in v)
                                    return v[t];
                                switch (t) {
                                case "keys":
                                case "values":
                                    return function() {
                                        return new n(this, t)
                                    }
                                }
                                return function() {
                                    return new n(this, t)
                                }
                            },
                            l = e + " Iterator",
                            h = "values" == o,
                            d = !1,
                            v = t.prototype,
                            p = v[dr] || v["@@iterator"] || o && v[o],
                            y = p || s(o),
                            g = o ? h ? s("entries") : y : void 0,
                            m = "Array" == e && v.entries || p;
                        if (m && (f = hr(m.call(new t))) !== Object.prototype && f.next && (Ve(f, l, !0), Te || "function" == typeof f[dr] || X(f, dr, pr)), h && p && "values" !== p.name && (d = !0, y = function() {
                            return p.call(this)
                        }), Te && !u || !vr && !d && v[dr] || X(v, dr, y), pn[e] = y, pn[l] = pr, o)
                            if (a = {
                                values: h ? y : s("values"),
                                keys: i ? y : s("keys"),
                                entries: g
                            }, u)
                                for (c in a)
                                    c in v || at(v, c, a[c]);
                            else
                                Qe(Qe.P + Qe.F * (vr || d), e, a);
                        return a
                    },
                    gr = Object.freeze({
                        default: yr,
                        __moduleExports: yr
                    }),
                    mr = or && rr || or,
                    br = ur && ir || ur,
                    _r = gr && yr || gr,
                    Or = _r(Array, "Array", (function(t, e) {
                        this._t = Zt(t),
                        this._i = 0,
                        this._k = e
                    }), (function() {
                        var t = this._t,
                            e = this._k,
                            n = this._i++;
                        return t && t.length > n ? br(0, "keys" == e ? n : "values" == e ? t[n] : [n, t[n]]) : (this._t = void 0, br(1))
                    }), "values");
                pn.Arguments = pn.Array,
                mr("keys"),
                mr("values"),
                mr("entries");
                var jr = ke("iterator"),
                    Er = !1;
                try {
                    [7][jr]().return = function() {
                        Er = !0
                    }
                } catch (t) {}
                var wr = function(t, e) {
                        if (!e && !Er)
                            return !1;
                        var n = !1;
                        try {
                            var r = [7],
                                o = r[jr]();
                            o.next = function() {
                                return {
                                    done: n = !0
                                }
                            },
                            r[jr] = function() {
                                return o
                            },
                            t(r)
                        } catch (t) {}
                        return n
                    },
                    xr = Object.freeze({
                        default: wr,
                        __moduleExports: wr
                    }),
                    Sr = [].copyWithin || function(t, e) {
                        var n = Se(this),
                            r = Nt(n.length),
                            o = te(t, r),
                            i = te(e, r),
                            u = arguments.length > 2 ? arguments[2] : void 0,
                            a = Math.min((void 0 === u ? r : te(u, r)) - i, r - o),
                            c = 1;
                        for (o > i && i + a > o && (c = -1, i += a - 1, o += a - 1); a-- > 0;)
                            i in n ? n[o] = n[i] : delete n[o],
                            o += c,
                            i += c;
                        return n
                    },
                    Mr = Object.freeze({
                        default: Sr,
                        __moduleExports: Sr
                    }),
                    zr = {}.propertyIsEnumerable,
                    Ar = {
                        f: zr
                    },
                    Pr = Object.freeze({
                        default: Ar,
                        __moduleExports: Ar,
                        f: zr
                    }),
                    kr = Pr && Ar || Pr,
                    Lr = Object.getOwnPropertyDescriptor,
                    Fr = w ? Lr : function(t, e) {
                        if (t = Zt(t), e = L(e, !0), k)
                            try {
                                return Lr(t, e)
                            } catch (t) {}
                        if ($(t, e))
                            return D(!kr.f.call(t, e), t[e])
                    },
                    Nr = {
                        f: Fr
                    },
                    Ir = Object.freeze({
                        default: Nr,
                        __moduleExports: Nr,
                        f: Fr
                    }),
                    Tr = bn && mn || bn,
                    Rr = Kn && Bn || Kn,
                    Cr = tr && Zn || tr,
                    Wr = xr && wr || xr,
                    Dr = Mr && Sr || Mr,
                    Br = Ir && Nr || Ir,
                    Kr = r((function(t) {
                        if (w) {
                            var e = Te,
                                n = b,
                                r = y,
                                o = Qe,
                                i = Re,
                                u = ct,
                                c = We,
                                f = D,
                                s = X,
                                h = Ce,
                                d = Pt,
                                v = Nt,
                                p = De,
                                g = te,
                                m = L,
                                _ = $,
                                O = Wn,
                                j = l,
                                E = Se,
                                x = Tr,
                                S = ar,
                                M = hr,
                                z = Be.f,
                                A = Rr,
                                P = Q,
                                k = ke,
                                F = Cr,
                                N = se,
                                I = tn,
                                T = Or,
                                R = pn,
                                C = Wr,
                                B = en,
                                K = Ke,
                                V = Dr,
                                U = W,
                                G = Br,
                                H = U.f,
                                J = G.f,
                                Y = n.RangeError,
                                q = n.TypeError,
                                Z = n.Uint8Array,
                                tt = Array.prototype,
                                et = Ze.ArrayBuffer,
                                nt = Ze.DataView,
                                rt = F(0),
                                ot = F(2),
                                it = F(3),
                                ut = F(4),
                                at = F(5),
                                ft = F(6),
                                st = N(!0),
                                lt = N(!1),
                                ht = T.values,
                                dt = T.keys,
                                vt = T.entries,
                                pt = tt.lastIndexOf,
                                yt = tt.reduce,
                                gt = tt.reduceRight,
                                mt = tt.join,
                                bt = tt.sort,
                                _t = tt.slice,
                                Ot = tt.toString,
                                jt = tt.toLocaleString,
                                Et = k("iterator"),
                                wt = k("toStringTag"),
                                xt = P("typed_constructor"),
                                St = P("def_constructor"),
                                Mt = i.CONSTR,
                                zt = i.TYPED,
                                At = i.VIEW,
                                kt = F(1, (function(t, e) {
                                    return Rt(I(t, t[St]), e)
                                })),
                                Lt = r((function() {
                                    return 1 === new Z(new Uint16Array([1]).buffer)[0]
                                })),
                                Ft = !!Z && !!Z.prototype.set && r((function() {
                                    new Z(1).set({})
                                })),
                                It = function(t, e) {
                                    var n = d(t);
                                    if (0 > n || n % e)
                                        throw Y("Wrong offset!");
                                    return n
                                },
                                Tt = function(t) {
                                    if (j(t) && zt in t)
                                        return t;
                                    throw q(t + " is not a typed array!")
                                },
                                Rt = function(t, e) {
                                    if (!j(t) || !(xt in t))
                                        throw q("It is not a typed array constructor!");
                                    return new t(e)
                                },
                                Ct = function(t, e) {
                                    return Wt(I(t, t[St]), e)
                                },
                                Wt = function(t, e) {
                                    for (var n = 0, r = e.length, o = Rt(t, r); r > n;)
                                        o[n] = e[n++];
                                    return o
                                },
                                Dt = function(t, e, n) {
                                    H(t, e, {
                                        get: function() {
                                            return this._d[n]
                                        }
                                    })
                                },
                                Bt = function(t) {
                                    var e,
                                        n,
                                        r,
                                        o,
                                        i,
                                        a,
                                        c = E(t),
                                        f = arguments.length,
                                        s = f > 1 ? arguments[1] : void 0,
                                        l = void 0 !== s,
                                        h = A(c);
                                    if (void 0 != h && !x(h)) {
                                        for (a = h.call(c), r = [], e = 0; !(i = a.next()).done; e++)
                                            r.push(i.value);
                                        c = r
                                    }
                                    for (l && f > 2 && (s = u(s, arguments[2], 2)), e = 0, n = v(c.length), o = Rt(this, n); n > e; e++)
                                        o[e] = l ? s(c[e], e) : c[e];
                                    return o
                                },
                                Kt = function() {
                                    for (var t = 0, e = arguments.length, n = Rt(this, e); e > t;)
                                        n[t] = arguments[t++];
                                    return n
                                },
                                Vt = !!Z && r((function() {
                                    jt.call(new Z(1))
                                })),
                                Ut = function() {
                                    return jt.apply(Vt ? _t.call(Tt(this)) : Tt(this), arguments)
                                },
                                Gt = {
                                    copyWithin: function(t, e) {
                                        return V.call(Tt(this), t, e, arguments.length > 2 ? arguments[2] : void 0)
                                    },
                                    every: function(t) {
                                        return ut(Tt(this), t, arguments.length > 1 ? arguments[1] : void 0)
                                    },
                                    fill: function(t) {
                                        return K.apply(Tt(this), arguments)
                                    },
                                    filter: function(t) {
                                        return Ct(this, ot(Tt(this), t, arguments.length > 1 ? arguments[1] : void 0))
                                    },
                                    find: function(t) {
                                        return at(Tt(this), t, arguments.length > 1 ? arguments[1] : void 0)
                                    },
                                    findIndex: function(t) {
                                        return ft(Tt(this), t, arguments.length > 1 ? arguments[1] : void 0)
                                    },
                                    forEach: function(t) {
                                        rt(Tt(this), t, arguments.length > 1 ? arguments[1] : void 0)
                                    },
                                    indexOf: function(t) {
                                        return lt(Tt(this), t, arguments.length > 1 ? arguments[1] : void 0)
                                    },
                                    includes: function(t) {
                                        return st(Tt(this), t, arguments.length > 1 ? arguments[1] : void 0)
                                    },
                                    join: function(t) {
                                        return mt.apply(Tt(this), arguments)
                                    },
                                    lastIndexOf: function(t) {
                                        return pt.apply(Tt(this), arguments)
                                    },
                                    map: function(t) {
                                        return kt(Tt(this), t, arguments.length > 1 ? arguments[1] : void 0)
                                    },
                                    reduce: function(t) {
                                        return yt.apply(Tt(this), arguments)
                                    },
                                    reduceRight: function(t) {
                                        return gt.apply(Tt(this), arguments)
                                    },
                                    reverse: function() {
                                        for (var t, e = Tt(this).length, n = Math.floor(e / 2), r = 0; n > r;)
                                            t = this[r],
                                            this[r++] = this[--e],
                                            this[e] = t;
                                        return this
                                    },
                                    some: function(t) {
                                        return it(Tt(this), t, arguments.length > 1 ? arguments[1] : void 0)
                                    },
                                    sort: function(t) {
                                        return bt.call(Tt(this), t)
                                    },
                                    subarray: function(t, e) {
                                        var n = Tt(this),
                                            r = n.length,
                                            o = g(t, r);
                                        return new (I(n, n[St]))(n.buffer, n.byteOffset + o * n.BYTES_PER_ELEMENT, v((void 0 === e ? r : g(e, r)) - o))
                                    }
                                },
                                Ht = function(t, e) {
                                    return Ct(this, _t.call(Tt(this), t, e))
                                },
                                Jt = function(t) {
                                    Tt(this);
                                    var e = It(arguments[1], 1),
                                        n = this.length,
                                        r = E(t),
                                        o = v(r.length),
                                        i = 0;
                                    if (o + e > n)
                                        throw Y("Wrong length!");
                                    for (; o > i;)
                                        this[e + i] = r[i++]
                                },
                                Yt = {
                                    entries: function() {
                                        return vt.call(Tt(this))
                                    },
                                    keys: function() {
                                        return dt.call(Tt(this))
                                    },
                                    values: function() {
                                        return ht.call(Tt(this))
                                    }
                                },
                                qt = function(t, e) {
                                    return j(t) && t[zt] && "symbol" != a(e) && e in t && +e + "" == e + ""
                                },
                                Xt = function(t, e) {
                                    return qt(t, e = m(e, !0)) ? f(2, t[e]) : J(t, e)
                                },
                                $t = function(t, e, n) {
                                    return !(qt(t, e = m(e, !0)) && j(n) && _(n, "value")) || _(n, "get") || _(n, "set") || n.configurable || _(n, "writable") && !n.writable || _(n, "enumerable") && !n.enumerable ? H(t, e, n) : (t[e] = n.value, t)
                                };
                            Mt || (G.f = Xt, U.f = $t),
                            o(o.S + o.F * !Mt, "Object", {
                                getOwnPropertyDescriptor: Xt,
                                defineProperty: $t
                            }),
                            r((function() {
                                Ot.call({})
                            })) && (Ot = jt = function() {
                                return mt.call(this)
                            });
                            var Qt = h({}, Gt);
                            h(Qt, Yt),
                            s(Qt, Et, Yt.values),
                            h(Qt, {
                                slice: Ht,
                                set: Jt,
                                constructor: function() {},
                                toString: Ot,
                                toLocaleString: Ut
                            }),
                            Dt(Qt, "buffer", "b"),
                            Dt(Qt, "byteOffset", "o"),
                            Dt(Qt, "byteLength", "l"),
                            Dt(Qt, "length", "e"),
                            H(Qt, wt, {
                                get: function() {
                                    return this[zt]
                                }
                            }),
                            t.exports = function(t, u, a, f) {
                                var l = t + ((f = !!f) ? "Clamped" : "") + "Array",
                                    h = "get" + t,
                                    d = "set" + t,
                                    y = n[l],
                                    g = y || {},
                                    m = y && M(y),
                                    b = {},
                                    _ = y && y.prototype,
                                    E = function(t, e) {
                                        H(t, e, {
                                            get: function() {
                                                return function(t, e) {
                                                    var n = t._d;
                                                    return n.v[h](e * u + n.o, Lt)
                                                }(this, e)
                                            },
                                            set: function(t) {
                                                return function(t, e, n) {
                                                    var r = t._d;
                                                    f && (n = 0 > (n = Math.round(n)) ? 0 : n > 255 ? 255 : 255 & n),
                                                    r.v[d](e * u + r.o, n, Lt)
                                                }(this, e, t)
                                            },
                                            enumerable: !0
                                        })
                                    };
                                y && i.ABV ? r((function() {
                                    y(1)
                                })) && r((function() {
                                    new y(-1)
                                })) && C((function(t) {
                                    new y,
                                    new y(null),
                                    new y(1.5),
                                    new y(t)
                                }), !0) || (y = a((function(t, e, n, r) {
                                    var o;
                                    return c(t, y, l), j(e) ? e instanceof et || "ArrayBuffer" == (o = O(e)) || "SharedArrayBuffer" == o ? void 0 !== r ? new g(e, It(n, u), r) : void 0 !== n ? new g(e, It(n, u)) : new g(e) : zt in e ? Wt(y, e) : Bt.call(y, e) : new g(p(e))
                                })), rt(m !== Function.prototype ? z(g).concat(z(m)) : z(g), (function(t) {
                                    t in y || s(y, t, g[t])
                                })), y.prototype = _, e || (_.constructor = y)) : (y = a((function(t, e, n, r) {
                                    c(t, y, l, "_d");
                                    var o,
                                        i,
                                        a,
                                        f,
                                        h = 0,
                                        d = 0;
                                    if (j(e)) {
                                        if (!(e instanceof et || "ArrayBuffer" == (f = O(e)) || "SharedArrayBuffer" == f))
                                            return zt in e ? Wt(y, e) : Bt.call(y, e);
                                        o = e,
                                        d = It(n, u);
                                        var g = e.byteLength;
                                        if (void 0 === r) {
                                            if (g % u)
                                                throw Y("Wrong length!");
                                            if (0 > (i = g - d))
                                                throw Y("Wrong length!")
                                        } else if ((i = v(r) * u) + d > g)
                                            throw Y("Wrong length!");
                                        a = i / u
                                    } else
                                        a = p(e),
                                        o = new et(i = a * u);
                                    for (s(t, "_d", {
                                        b: o,
                                        o: d,
                                        l: i,
                                        e: a,
                                        v: new nt(o)
                                    }); a > h;)
                                        E(t, h++)
                                })), _ = y.prototype = S(Qt), s(_, "constructor", y));
                                var w = _[Et],
                                    x = !!w && ("values" == w.name || void 0 == w.name),
                                    A = Yt.values;
                                s(y, xt, !0),
                                s(_, zt, l),
                                s(_, At, !0),
                                s(_, St, y),
                                (f ? new y(1)[wt] == l : wt in _) || H(_, wt, {
                                    get: function() {
                                        return l
                                    }
                                }),
                                b[l] = y,
                                o(o.G + o.W + o.F * (y != g), b),
                                o(o.S, l, {
                                    BYTES_PER_ELEMENT: u
                                }),
                                o(o.S + o.F * r((function() {
                                    g.of.call(y, 1)
                                })), l, {
                                    from: Bt,
                                    of: Kt
                                }),
                                "BYTES_PER_ELEMENT" in _ || s(_, "BYTES_PER_ELEMENT", u),
                                o(o.P, l, Gt),
                                B(l),
                                o(o.P + o.F * Ft, l, {
                                    set: Jt
                                }),
                                o(o.P + o.F * !x, l, Yt),
                                e || _.toString == Ot || (_.toString = Ot),
                                o(o.P + o.F * r((function() {
                                    new y(1).slice()
                                })), l, {
                                    slice: Ht
                                }),
                                o(o.P + o.F * (r((function() {
                                    return [1, 2].toLocaleString() != new y([1, 2]).toLocaleString()
                                })) || !r((function() {
                                    _.toLocaleString.call([1, 2])
                                }))), l, {
                                    toLocaleString: Ut
                                }),
                                R[l] = x ? w : A,
                                e || x || s(_, Et, A)
                            }
                        } else
                            t.exports = function() {}
                    })),
                    Vr = Object.freeze({
                        default: Kr,
                        __moduleExports: Kr
                    }),
                    Ur = Vr && Kr || Vr;
                Ur("Int8", 1, (function(t) {
                    return function(e, n, r) {
                        return t(this, e, n, r)
                    }
                })),
                Ur("Uint8", 1, (function(t) {
                    return function(e, n, r) {
                        return t(this, e, n, r)
                    }
                })),
                Ur("Uint8", 1, (function(t) {
                    return function(e, n, r) {
                        return t(this, e, n, r)
                    }
                }), !0),
                Ur("Int16", 2, (function(t) {
                    return function(e, n, r) {
                        return t(this, e, n, r)
                    }
                })),
                Ur("Uint16", 2, (function(t) {
                    return function(e, n, r) {
                        return t(this, e, n, r)
                    }
                })),
                Ur("Int32", 4, (function(t) {
                    return function(e, n, r) {
                        return t(this, e, n, r)
                    }
                })),
                Ur("Uint32", 4, (function(t) {
                    return function(e, n, r) {
                        return t(this, e, n, r)
                    }
                })),
                Ur("Float32", 4, (function(t) {
                    return function(e, n, r) {
                        return t(this, e, n, r)
                    }
                })),
                Ur("Float64", 8, (function(t) {
                    return function(e, n, r) {
                        return t(this, e, n, r)
                    }
                }));
                var Gr = function(t, e, n, r) {
                        try {
                            return r ? e(P(n)[0], n[1]) : e(n)
                        } catch (e) {
                            var o = t.return;
                            throw void 0 !== o && P(o.call(t)), e
                        }
                    },
                    Hr = Object.freeze({
                        default: Gr,
                        __moduleExports: Gr
                    }),
                    Jr = Hr && Gr || Hr,
                    Yr = r((function(t) {
                        var e = {},
                            n = {},
                            r = t.exports = function(t, r, o, i, u) {
                                var a,
                                    c,
                                    f,
                                    s,
                                    l = u ? function() {
                                        return t
                                    } : Rr(t),
                                    h = ct(o, i, r ? 2 : 1),
                                    d = 0;
                                if ("function" != typeof l)
                                    throw TypeError(t + " is not iterable!");
                                if (Tr(l)) {
                                    for (a = Nt(t.length); a > d; d++)
                                        if ((s = r ? h(P(c = t[d])[0], c[1]) : h(t[d])) === e || s === n)
                                            return s
                                } else
                                    for (f = l.call(t); !(c = f.next()).done;)
                                        if ((s = Jr(f, h, c.value, r)) === e || s === n)
                                            return s
                            };
                        r.BREAK = e,
                        r.RETURN = n
                    })),
                    qr = Object.freeze({
                        default: Yr,
                        __moduleExports: Yr
                    }),
                    Xr = r((function(t) {
                        var e = Q("meta"),
                            n = W.f,
                            r = 0,
                            o = Object.isExtensible || function() {
                                return !0
                            },
                            i = !y((function() {
                                return o(Object.preventExtensions({}))
                            })),
                            u = function(t) {
                                n(t, e, {
                                    value: {
                                        i: "O" + ++r,
                                        w: {}
                                    }
                                })
                            },
                            c = t.exports = {
                                KEY: e,
                                NEED: !1,
                                fastKey: function(t, n) {
                                    if (!l(t))
                                        return "symbol" == a(t) ? t : ("string" == typeof t ? "S" : "P") + t;
                                    if (!$(t, e)) {
                                        if (!o(t))
                                            return "F";
                                        if (!n)
                                            return "E";
                                        u(t)
                                    }
                                    return t[e].i
                                },
                                getWeak: function(t, n) {
                                    if (!$(t, e)) {
                                        if (!o(t))
                                            return !0;
                                        if (!n)
                                            return !1;
                                        u(t)
                                    }
                                    return t[e].w
                                },
                                onFreeze: function(t) {
                                    return i && c.NEED && o(t) && !$(t, e) && u(t), t
                                }
                            }
                    })),
                    $r = Object.freeze({
                        default: Xr,
                        __moduleExports: Xr,
                        KEY: Xr.KEY,
                        NEED: Xr.NEED,
                        fastKey: Xr.fastKey,
                        getWeak: Xr.getWeak,
                        onFreeze: Xr.onFreeze
                    }),
                    Qr = function(t, e) {
                        if (!l(t) || t._t !== e)
                            throw TypeError("Incompatible receiver, " + e + " required!");
                        return t
                    },
                    Zr = Object.freeze({
                        default: Qr,
                        __moduleExports: Qr
                    }),
                    to = qr && Yr || qr,
                    eo = $r && Xr || $r,
                    no = Zr && Qr || Zr,
                    ro = W.f,
                    oo = eo.fastKey,
                    io = w ? "_s" : "size",
                    uo = function(t, e) {
                        var n,
                            r = oo(e);
                        if ("F" !== r)
                            return t._i[r];
                        for (n = t._f; n; n = n.n)
                            if (n.k == e)
                                return n
                    },
                    ao = {
                        getConstructor: function(t, e, n, r) {
                            var o = t((function(t, i) {
                                We(t, o, e, "_i"),
                                t._t = e,
                                t._i = ar(null),
                                t._f = void 0,
                                t._l = void 0,
                                t[io] = 0,
                                void 0 != i && to(i, n, t[r], t)
                            }));
                            return Ce(o.prototype, {
                                clear: function() {
                                    for (var t = no(this, e), n = t._i, r = t._f; r; r = r.n)
                                        r.r = !0,
                                        r.p && (r.p = r.p.n = void 0),
                                        delete n[r.i];
                                    t._f = t._l = void 0,
                                    t[io] = 0
                                },
                                delete: function(t) {
                                    var n = no(this, e),
                                        r = uo(n, t);
                                    if (r) {
                                        var o = r.n,
                                            i = r.p;
                                        delete n._i[r.i],
                                        r.r = !0,
                                        i && (i.n = o),
                                        o && (o.p = i),
                                        n._f == r && (n._f = o),
                                        n._l == r && (n._l = i),
                                        n[io]--
                                    }
                                    return !!r
                                },
                                forEach: function(t) {
                                    no(this, e);
                                    for (var n, r = ct(t, arguments.length > 1 ? arguments[1] : void 0, 3); n = n ? n.n : this._f;)
                                        for (r(n.v, n.k, this); n && n.r;)
                                            n = n.p
                                },
                                has: function(t) {
                                    return !!uo(no(this, e), t)
                                }
                            }), w && ro(o.prototype, "size", {
                                get: function() {
                                    return no(this, e)[io]
                                }
                            }), o
                        },
                        def: function(t, e, n) {
                            var r,
                                o,
                                i = uo(t, e);
                            return i ? i.v = n : (t._l = i = {
                                i: o = oo(e, !0),
                                k: e,
                                v: n,
                                p: r = t._l,
                                n: void 0,
                                r: !1
                            }, t._f || (t._f = i), r && (r.n = i), t[io]++, "F" !== o && (t._i[o] = i)), t
                        },
                        getEntry: uo,
                        setStrong: function(t, e, n) {
                            _r(t, e, (function(t, n) {
                                this._t = no(t, e),
                                this._k = n,
                                this._l = void 0
                            }), (function() {
                                for (var t = this._k, e = this._l; e && e.r;)
                                    e = e.p;
                                return this._t && (this._l = e = e ? e.n : this._t._f) ? br(0, "keys" == t ? e.k : "values" == t ? e.v : [e.k, e.v]) : (this._t = void 0, br(1))
                            }), n ? "entries" : "values", !n, !0),
                            en(e)
                        }
                    },
                    co = Object.freeze({
                        default: ao,
                        __moduleExports: ao,
                        getConstructor: ao.getConstructor,
                        def: ao.def,
                        getEntry: ao.getEntry,
                        setStrong: ao.setStrong
                    }),
                    fo = function(t, e) {
                        if (P(t), !l(e) && null !== e)
                            throw TypeError(e + ": can't set as prototype!")
                    },
                    so = {
                        set: Object.setPrototypeOf || ("__proto__" in {} ? function(t, e, n) {
                            try {
                                (n = ct(Function.call, Br.f(Object.prototype, "__proto__").set, 2))(t, []),
                                e = !(t instanceof Array)
                            } catch (t) {
                                e = !0
                            }
                            return function(t, r) {
                                return fo(t, r), e ? t.__proto__ = r : n(t, r), t
                            }
                        }({}, !1) : void 0),
                        check: fo
                    },
                    lo = Object.freeze({
                        default: so,
                        __moduleExports: so,
                        set: so.set,
                        check: so.check
                    }),
                    ho = lo && so || lo,
                    vo = ho.set,
                    po = function(t, e, n) {
                        var r,
                            o = e.constructor;
                        return o !== n && "function" == typeof o && (r = o.prototype) !== n.prototype && l(r) && vo && vo(t, r), t
                    },
                    yo = Object.freeze({
                        default: po,
                        __moduleExports: po
                    }),
                    go = yo && po || yo,
                    mo = function(t, e, n, r, o, i) {
                        var u = b[t],
                            a = u,
                            c = o ? "set" : "add",
                            f = a && a.prototype,
                            s = {},
                            h = function(t) {
                                var e = f[t];
                                at(f, t, "delete" == t || "has" == t ? function(t) {
                                    return !(i && !l(t)) && e.call(this, 0 === t ? 0 : t)
                                } : "get" == t ? function(t) {
                                    return i && !l(t) ? void 0 : e.call(this, 0 === t ? 0 : t)
                                } : "add" == t ? function(t) {
                                    return e.call(this, 0 === t ? 0 : t), this
                                } : function(t, n) {
                                    return e.call(this, 0 === t ? 0 : t, n), this
                                })
                            };
                        if ("function" == typeof a && (i || f.forEach && !y((function() {
                            (new a).entries().next()
                        })))) {
                            var d = new a,
                                v = d[c](i ? {} : -0, 1) != d,
                                p = y((function() {
                                    d.has(1)
                                })),
                                g = Wr((function(t) {
                                    new a(t)
                                })),
                                m = !i && y((function() {
                                    for (var t = new a, e = 5; e--;)
                                        t[c](e, e);
                                    return !t.has(-0)
                                }));
                            g || ((a = e((function(e, n) {
                                We(e, a, t);
                                var r = go(new u, e, a);
                                return void 0 != n && to(n, o, r[c], r), r
                            }))).prototype = f, f.constructor = a),
                            (p || m) && (h("delete"), h("has"), o && h("get")),
                            (m || v) && h(c),
                            i && f.clear && delete f.clear
                        } else
                            a = r.getConstructor(e, t, o, c),
                            Ce(a.prototype, n),
                            eo.NEED = !0;
                        return Ve(a, t), s[t] = a, Qe(Qe.G + Qe.W + Qe.F * (a != u), s), i || r.setStrong(a, t, o), a
                    },
                    bo = Object.freeze({
                        default: mo,
                        __moduleExports: mo
                    }),
                    _o = co && ao || co,
                    Oo = bo && mo || bo,
                    jo = (Oo("Map", (function(t) {
                        return function() {
                            return t(this, arguments.length > 0 ? arguments[0] : void 0)
                        }
                    }), {
                        get: function(t) {
                            var e = _o.getEntry(no(this, "Map"), t);
                            return e && e.v
                        },
                        set: function(t, e) {
                            return _o.def(no(this, "Map"), 0 === t ? 0 : t, e)
                        }
                    }, _o, !0), Oo("Set", (function(t) {
                        return function() {
                            return t(this, arguments.length > 0 ? arguments[0] : void 0)
                        }
                    }), {
                        add: function(t) {
                            return _o.def(no(this, "Set"), t = 0 === t ? 0 : t, t)
                        }
                    }, _o), Object.getOwnPropertySymbols),
                    Eo = {
                        f: jo
                    },
                    wo = Object.freeze({
                        default: Eo,
                        __moduleExports: Eo,
                        f: jo
                    }),
                    xo = wo && Eo || wo,
                    So = Object.assign,
                    Mo = !So || y((function() {
                        var t = {},
                            e = {},
                            n = Symbol(),
                            r = "abcdefghijklmnopqrst";
                        return t[n] = 7, r.split("").forEach((function(t) {
                            e[t] = t
                        })), 7 != So({}, t)[n] || Object.keys(So({}, e)).join("") != r
                    })) ? function(t, e) {
                        for (var n = Se(t), r = arguments.length, o = 1, i = xo.f, u = kr.f; r > o;)
                            for (var a, c = Gt(arguments[o++]), f = i ? jn(c).concat(i(c)) : jn(c), s = f.length, l = 0; s > l;)
                                u.call(c, a = f[l++]) && (n[a] = c[a]);
                        return n
                    } : So,
                    zo = Object.freeze({
                        default: Mo,
                        __moduleExports: Mo
                    }),
                    Ao = eo.getWeak,
                    Po = Cr(5),
                    ko = Cr(6),
                    Lo = 0,
                    Fo = function(t) {
                        return t._l || (t._l = new No)
                    },
                    No = function() {
                        this.a = []
                    },
                    Io = function(t, e) {
                        return Po(t.a, (function(t) {
                            return t[0] === e
                        }))
                    };
                No.prototype = {
                    get: function(t) {
                        var e = Io(this, t);
                        if (e)
                            return e[1]
                    },
                    has: function(t) {
                        return !!Io(this, t)
                    },
                    set: function(t, e) {
                        var n = Io(this, t);
                        n ? n[1] = e : this.a.push([t, e])
                    },
                    delete: function(t) {
                        var e = ko(this.a, (function(e) {
                            return e[0] === t
                        }));
                        return ~e && this.a.splice(e, 1), !!~e
                    }
                };
                var To = {
                        getConstructor: function(t, e, n, r) {
                            var o = t((function(t, i) {
                                We(t, o, e, "_i"),
                                t._t = e,
                                t._i = Lo++,
                                t._l = void 0,
                                void 0 != i && to(i, n, t[r], t)
                            }));
                            return Ce(o.prototype, {
                                delete: function(t) {
                                    if (!l(t))
                                        return !1;
                                    var n = Ao(t);
                                    return !0 === n ? Fo(no(this, e)).delete(t) : n && $(n, this._i) && delete n[this._i]
                                },
                                has: function(t) {
                                    if (!l(t))
                                        return !1;
                                    var n = Ao(t);
                                    return !0 === n ? Fo(no(this, e)).has(t) : n && $(n, this._i)
                                }
                            }), o
                        },
                        def: function(t, e, n) {
                            var r = Ao(P(e), !0);
                            return !0 === r ? Fo(t).set(e, n) : r[t._i] = n, t
                        },
                        ufstore: Fo
                    },
                    Ro = Object.freeze({
                        default: To,
                        __moduleExports: To,
                        getConstructor: To.getConstructor,
                        def: To.def,
                        ufstore: To.ufstore
                    }),
                    Co = zo && Mo || zo,
                    Wo = Ro && To || Ro;
                r((function(t) {
                    var e,
                        n = Cr(0),
                        r = eo.getWeak,
                        o = Object.isExtensible,
                        i = Wo.ufstore,
                        u = {},
                        a = function(t) {
                            return function() {
                                return t(this, arguments.length > 0 ? arguments[0] : void 0)
                            }
                        },
                        c = {
                            get: function(t) {
                                if (l(t)) {
                                    var e = r(t);
                                    return !0 === e ? i(no(this, "WeakMap")).get(t) : e ? e[this._i] : void 0
                                }
                            },
                            set: function(t, e) {
                                return Wo.def(no(this, "WeakMap"), t, e)
                            }
                        },
                        f = t.exports = Oo("WeakMap", a, c, Wo, !0, !0);
                    y((function() {
                        return 7 != (new f).set((Object.freeze || Object)(u), 7).get(u)
                    })) && (e = Wo.getConstructor(a, "WeakMap"), Co(e.prototype, c), eo.NEED = !0, n(["delete", "has", "get", "set"], (function(t) {
                        var n = f.prototype,
                            r = n[t];
                        at(n, t, (function(n, i) {
                            if (l(n) && !o(n)) {
                                this._f || (this._f = new e);
                                var u = this._f[t](n, i);
                                return "set" == t ? this : u
                            }
                            return r.call(this, n, i)
                        }))
                    })))
                })),
                Oo("WeakSet", (function(t) {
                    return function() {
                        return t(this, arguments.length > 0 ? arguments[0] : void 0)
                    }
                }), {
                    add: function(t) {
                        return Wo.def(no(this, "WeakSet"), t, !0)
                    }
                }, Wo, !1, !0);
                var Do = (b.Reflect || {}).apply,
                    Bo = Function.apply;
                Qe(Qe.S + Qe.F * !y((function() {
                    Do((function() {}))
                })), "Reflect", {
                    apply: function(t, e, n) {
                        var r = ot(t),
                            o = P(n);
                        return Do ? Do(r, e, o) : Bo.call(r, e, o)
                    }
                });
                var Ko = function(t, e, n) {
                        var r = void 0 === n;
                        switch (e.length) {
                        case 0:
                            return r ? t() : t.call(n);
                        case 1:
                            return r ? t(e[0]) : t.call(n, e[0]);
                        case 2:
                            return r ? t(e[0], e[1]) : t.call(n, e[0], e[1]);
                        case 3:
                            return r ? t(e[0], e[1], e[2]) : t.call(n, e[0], e[1], e[2]);
                        case 4:
                            return r ? t(e[0], e[1], e[2], e[3]) : t.call(n, e[0], e[1], e[2], e[3])
                        }
                        return t.apply(n, e)
                    },
                    Vo = Object.freeze({
                        default: Ko,
                        __moduleExports: Ko
                    }),
                    Uo = Vo && Ko || Vo,
                    Go = [].slice,
                    Ho = {},
                    Jo = Function.bind || function(t) {
                        var e = ot(this),
                            n = Go.call(arguments, 1),
                            r = function r() {
                                var o = n.concat(Go.call(arguments));
                                return this instanceof r ? function(t, e, n) {
                                    if (!(e in Ho)) {
                                        for (var r = [], o = 0; e > o; o++)
                                            r[o] = "a[" + o + "]";
                                        Ho[e] = Function("F,a", "return new F(" + r.join(",") + ")")
                                    }
                                    return Ho[e](t, n)
                                }(e, o.length, o) : Uo(e, o, t)
                            };
                        return l(e.prototype) && (r.prototype = e.prototype), r
                    },
                    Yo = Object.freeze({
                        default: Jo,
                        __moduleExports: Jo
                    }),
                    qo = Yo && Jo || Yo,
                    Xo = (b.Reflect || {}).construct,
                    $o = y((function() {
                        function t() {}
                        return !(Xo((function() {}), [], t) instanceof t)
                    })),
                    Qo = !y((function() {
                        Xo((function() {}))
                    }));
                Qe(Qe.S + Qe.F * ($o || Qo), "Reflect", {
                    construct: function(t, e) {
                        ot(t),
                        P(e);
                        var n = 3 > arguments.length ? t : ot(arguments[2]);
                        if (Qo && !$o)
                            return Xo(t, e, n);
                        if (t == n) {
                            switch (e.length) {
                            case 0:
                                return new t;
                            case 1:
                                return new t(e[0]);
                            case 2:
                                return new t(e[0], e[1]);
                            case 3:
                                return new t(e[0], e[1], e[2]);
                            case 4:
                                return new t(e[0], e[1], e[2], e[3])
                            }
                            var r = [null];
                            return r.push.apply(r, e), new (qo.apply(t, r))
                        }
                        var o = n.prototype,
                            i = ar(l(o) ? o : Object.prototype),
                            u = Function.apply.call(t, i, e);
                        return l(u) ? u : i
                    }
                }),
                Qe(Qe.S + Qe.F * y((function() {
                    Reflect.defineProperty(W.f({}, 1, {
                        value: 1
                    }), 1, {
                        value: 2
                    })
                })), "Reflect", {
                    defineProperty: function(t, e, n) {
                        P(t),
                        e = L(e, !0),
                        P(n);
                        try {
                            return W.f(t, e, n), !0
                        } catch (t) {
                            return !1
                        }
                    }
                });
                var Zo = Br.f;
                Qe(Qe.S, "Reflect", {
                    deleteProperty: function(t, e) {
                        var n = Zo(P(t), e);
                        return !(n && !n.configurable) && delete t[e]
                    }
                }),
                Qe(Qe.S, "Reflect", {
                    get: function t(e, n) {
                        var r,
                            o,
                            i = 3 > arguments.length ? e : arguments[2];
                        return P(e) === i ? e[n] : (r = Br.f(e, n)) ? $(r, "value") ? r.value : void 0 !== r.get ? r.get.call(i) : void 0 : l(o = hr(e)) ? t(o, n, i) : void 0
                    }
                }),
                Qe(Qe.S, "Reflect", {
                    getOwnPropertyDescriptor: function(t, e) {
                        return Br.f(P(t), e)
                    }
                }),
                Qe(Qe.S, "Reflect", {
                    getPrototypeOf: function(t) {
                        return hr(P(t))
                    }
                }),
                Qe(Qe.S, "Reflect", {
                    has: function(t, e) {
                        return e in t
                    }
                });
                var ti = Object.isExtensible;
                Qe(Qe.S, "Reflect", {
                    isExtensible: function(t) {
                        return P(t), !ti || ti(t)
                    }
                });
                var ei = b.Reflect,
                    ni = ei && ei.ownKeys || function(t) {
                        var e = Be.f(P(t)),
                            n = xo.f;
                        return n ? e.concat(n(t)) : e
                    },
                    ri = Object.freeze({
                        default: ni,
                        __moduleExports: ni
                    }),
                    oi = ri && ni || ri;
                Qe(Qe.S, "Reflect", {
                    ownKeys: oi
                });
                var ii = Object.preventExtensions;
                Qe(Qe.S, "Reflect", {
                    preventExtensions: function(t) {
                        P(t);
                        try {
                            return ii && ii(t), !0
                        } catch (t) {
                            return !1
                        }
                    }
                }),
                Qe(Qe.S, "Reflect", {
                    set: function t(e, n, r) {
                        var o,
                            i,
                            u = 4 > arguments.length ? e : arguments[3],
                            a = Br.f(P(e), n);
                        if (!a) {
                            if (l(i = hr(e)))
                                return t(i, n, r, u);
                            a = D(0)
                        }
                        if ($(a, "value")) {
                            if (!1 === a.writable || !l(u))
                                return !1;
                            if (o = Br.f(u, n)) {
                                if (o.get || o.set || !1 === o.writable)
                                    return !1;
                                o.value = r,
                                W.f(u, n, o)
                            } else
                                W.f(u, n, D(0, r));
                            return !0
                        }
                        return void 0 !== a.set && (a.set.call(u, r), !0)
                    }
                }),
                ho && Qe(Qe.S, "Reflect", {
                    setPrototypeOf: function(t, e) {
                        ho.check(t, e);
                        try {
                            return ho.set(t, e), !0
                        } catch (t) {
                            return !1
                        }
                    }
                });
                var ui,
                    ai,
                    ci,
                    fi = b.process,
                    si = b.setImmediate,
                    li = b.clearImmediate,
                    hi = b.MessageChannel,
                    di = b.Dispatch,
                    vi = 0,
                    pi = {},
                    yi = function() {
                        var t = +this;
                        if (pi.hasOwnProperty(t)) {
                            var e = pi[t];
                            delete pi[t],
                            e()
                        }
                    },
                    gi = function(t) {
                        yi.call(t.data)
                    };
                si && li || (si = function(t) {
                    for (var e = [], n = 1; arguments.length > n;)
                        e.push(arguments[n++]);
                    return pi[++vi] = function() {
                        Uo("function" == typeof t ? t : Function(t), e)
                    }, ui(vi), vi
                }, li = function(t) {
                    delete pi[t]
                }, "process" == Dt(fi) ? ui = function(t) {
                    fi.nextTick(ct(yi, t, 1))
                } : di && di.now ? ui = function(t) {
                    di.now(ct(yi, t, 1))
                } : hi ? (ci = (ai = new hi).port2, ai.port1.onmessage = gi, ui = ct(ci.postMessage, ci, 1)) : b.addEventListener && "function" == typeof postMessage && !b.importScripts ? (ui = function(t) {
                    b.postMessage(t + "", "*")
                }, b.addEventListener("message", gi, !1)) : ui = "onreadystatechange" in x("script") ? function(t) {
                    An.appendChild(x("script")).onreadystatechange = function() {
                        An.removeChild(this),
                        yi.call(t)
                    }
                } : function(t) {
                    setTimeout(ct(yi, t, 1), 0)
                });
                var mi,
                    bi,
                    _i,
                    Oi,
                    ji = {
                        set: si,
                        clear: li
                    },
                    Ei = Object.freeze({
                        default: ji,
                        __moduleExports: ji,
                        set: ji.set,
                        clear: ji.clear
                    }),
                    wi = Ei && ji || Ei,
                    xi = wi.set,
                    Si = b.MutationObserver || b.WebKitMutationObserver,
                    Mi = b.process,
                    zi = b.Promise,
                    Ai = "process" == Dt(Mi),
                    Pi = function() {
                        var t,
                            e,
                            n,
                            r = function() {
                                var r,
                                    o;
                                for (Ai && (r = Mi.domain) && r.exit(); t;) {
                                    o = t.fn,
                                    t = t.next;
                                    try {
                                        o()
                                    } catch (r) {
                                        throw t ? n() : e = void 0, r
                                    }
                                }
                                e = void 0,
                                r && r.enter()
                            };
                        if (Ai)
                            n = function() {
                                Mi.nextTick(r)
                            };
                        else if (!Si || b.navigator && b.navigator.standalone)
                            if (zi && zi.resolve) {
                                var o = zi.resolve();
                                n = function() {
                                    o.then(r)
                                }
                            } else
                                n = function() {
                                    xi.call(b, r)
                                };
                        else {
                            var i = !0,
                                u = document.createTextNode("");
                            new Si(r).observe(u, {
                                characterData: !0
                            }),
                            n = function() {
                                u.data = i = !i
                            }
                        }
                        return function(r) {
                            var o = {
                                fn: r,
                                next: void 0
                            };
                            e && (e.next = o),
                            t || (t = o, n()),
                            e = o
                        }
                    },
                    ki = Object.freeze({
                        default: Pi,
                        __moduleExports: Pi
                    }),
                    Li = function(t) {
                        return new function(t) {
                            var e,
                                n;
                            this.promise = new t((function(t, r) {
                                if (void 0 !== e || void 0 !== n)
                                    throw TypeError("Bad Promise constructor");
                                e = t,
                                n = r
                            })),
                            this.resolve = ot(e),
                            this.reject = ot(n)
                        }(t)
                    },
                    Fi = {
                        f: Li
                    },
                    Ni = Object.freeze({
                        default: Fi,
                        __moduleExports: Fi,
                        f: Li
                    }),
                    Ii = function(t) {
                        try {
                            return {
                                e: !1,
                                v: t()
                            }
                        } catch (t) {
                            return {
                                e: !0,
                                v: t
                            }
                        }
                    },
                    Ti = Object.freeze({
                        default: Ii,
                        __moduleExports: Ii
                    }),
                    Ri = Ni && Fi || Ni,
                    Ci = function(t, e) {
                        if (P(t), l(e) && e.constructor === t)
                            return e;
                        var n = Ri.f(t);
                        return (0, n.resolve)(e), n.promise
                    },
                    Wi = Object.freeze({
                        default: Ci,
                        __moduleExports: Ci
                    }),
                    Di = Ti && Ii || Ti,
                    Bi = Wi && Ci || Wi,
                    Ki = wi.set,
                    Vi = (ki && Pi || ki)(),
                    Ui = b.TypeError,
                    Gi = b.process,
                    Hi = b.Promise,
                    Ji = "process" == Wn(Gi),
                    Yi = function() {},
                    qi = bi = Ri.f,
                    Xi = !!function() {
                        try {
                            var t = Hi.resolve(1),
                                e = (t.constructor = {})[ke("species")] = function(t) {
                                    t(Yi, Yi)
                                };
                            return (Ji || "function" == typeof PromiseRejectionEvent) && t.then(Yi) instanceof e
                        } catch (t) {}
                    }(),
                    $i = function(t) {
                        var e;
                        return !(!l(t) || "function" != typeof (e = t.then)) && e
                    },
                    Qi = function(t, e) {
                        if (!t._n) {
                            t._n = !0;
                            var n = t._c;
                            Vi((function() {
                                for (var r = t._v, o = 1 == t._s, i = 0, u = function(e) {
                                        var n,
                                            i,
                                            u,
                                            a = o ? e.ok : e.fail,
                                            c = e.resolve,
                                            f = e.reject,
                                            s = e.domain;
                                        try {
                                            a ? (o || (2 == t._h && eu(t), t._h = 1), !0 === a ? n = r : (s && s.enter(), n = a(r), s && (s.exit(), u = !0)), n === e.promise ? f(Ui("Promise-chain cycle")) : (i = $i(n)) ? i.call(n, c, f) : c(n)) : f(r)
                                        } catch (t) {
                                            s && !u && s.exit(),
                                            f(t)
                                        }
                                    }; n.length > i;)
                                    u(n[i++]);
                                t._c = [],
                                t._n = !1,
                                e && !t._h && Zi(t)
                            }))
                        }
                    },
                    Zi = function(t) {
                        Ki.call(b, (function() {
                            var e,
                                n,
                                r,
                                o = t._v,
                                i = tu(t);
                            if (i && (e = Di((function() {
                                Ji ? Gi.emit("unhandledRejection", o, t) : (n = b.onunhandledrejection) ? n({
                                    promise: t,
                                    reason: o
                                }) : (r = b.console) && r.error && r.error("Unhandled promise rejection", o)
                            })), t._h = Ji || tu(t) ? 2 : 1), t._a = void 0, i && e.e)
                                throw e.v
                        }))
                    },
                    tu = function(t) {
                        return 1 !== t._h && 0 === (t._a || t._c).length
                    },
                    eu = function(t) {
                        Ki.call(b, (function() {
                            var e;
                            Ji ? Gi.emit("rejectionHandled", t) : (e = b.onrejectionhandled) && e({
                                promise: t,
                                reason: t._v
                            })
                        }))
                    },
                    nu = function(t) {
                        var e = this;
                        e._d || (e._d = !0, (e = e._w || e)._v = t, e._s = 2, e._a || (e._a = e._c.slice()), Qi(e, !0))
                    },
                    ru = function t(e) {
                        var n,
                            r = this;
                        if (!r._d) {
                            r._d = !0,
                            r = r._w || r;
                            try {
                                if (r === e)
                                    throw Ui("Promise can't be resolved itself");
                                (n = $i(e)) ? Vi((function() {
                                    var o = {
                                        _w: r,
                                        _d: !1
                                    };
                                    try {
                                        n.call(e, ct(t, o, 1), ct(nu, o, 1))
                                    } catch (e) {
                                        nu.call(o, e)
                                    }
                                })) : (r._v = e, r._s = 1, Qi(r, !1))
                            } catch (e) {
                                nu.call({
                                    _w: r,
                                    _d: !1
                                }, e)
                            }
                        }
                    };
                Xi || (Hi = function(t) {
                    We(this, Hi, "Promise", "_h"),
                    ot(t),
                    mi.call(this);
                    try {
                        t(ct(ru, this, 1), ct(nu, this, 1))
                    } catch (t) {
                        nu.call(this, t)
                    }
                }, (mi = function(t) {
                    this._c = [],
                    this._a = void 0,
                    this._s = 0,
                    this._d = !1,
                    this._v = void 0,
                    this._h = 0,
                    this._n = !1
                }).prototype = Ce(Hi.prototype, {
                    then: function(t, e) {
                        var n = qi(tn(this, Hi));
                        return n.ok = "function" != typeof t || t, n.fail = "function" == typeof e && e, n.domain = Ji ? Gi.domain : void 0, this._c.push(n), this._a && this._a.push(n), this._s && Qi(this, !1), n.promise
                    },
                    catch: function(t) {
                        return this.then(void 0, t)
                    }
                }), _i = function() {
                    var t = new mi;
                    this.promise = t,
                    this.resolve = ct(ru, t, 1),
                    this.reject = ct(nu, t, 1)
                }, Ri.f = qi = function(t) {
                    return t === Hi || t === Oi ? new _i(t) : bi(t)
                }),
                Qe(Qe.G + Qe.W + Qe.F * !Xi, {
                    Promise: Hi
                }),
                Ve(Hi, "Promise"),
                en("Promise"),
                Oi = Z.Promise,
                Qe(Qe.S + Qe.F * !Xi, "Promise", {
                    reject: function(t) {
                        var e = qi(this);
                        return (0, e.reject)(t), e.promise
                    }
                }),
                Qe(Qe.S + Qe.F * (Te || !Xi), "Promise", {
                    resolve: function(t) {
                        return Bi(Te && this === Oi ? Hi : this, t)
                    }
                }),
                Qe(Qe.S + Qe.F * !(Xi && Wr((function(t) {
                    Hi.all(t).catch(Yi)
                }))), "Promise", {
                    all: function(t) {
                        var e = this,
                            n = qi(e),
                            r = n.resolve,
                            o = n.reject,
                            i = Di((function() {
                                var n = [],
                                    i = 0,
                                    u = 1;
                                to(t, !1, (function(t) {
                                    var a = i++,
                                        c = !1;
                                    n.push(void 0),
                                    u++,
                                    e.resolve(t).then((function(t) {
                                        c || (c = !0, n[a] = t, --u || r(n))
                                    }), o)
                                })),
                                --u || r(n)
                            }));
                        return i.e && o(i.v), n.promise
                    },
                    race: function(t) {
                        var e = this,
                            n = qi(e),
                            r = n.reject,
                            o = Di((function() {
                                to(t, !1, (function(t) {
                                    e.resolve(t).then(n.resolve, r)
                                }))
                            }));
                        return o.e && r(o.v), n.promise
                    }
                });
                var ou = {
                        f: ke
                    },
                    iu = Object.freeze({
                        default: ou,
                        __moduleExports: ou,
                        f: ke
                    }),
                    uu = iu && ou || iu,
                    au = W.f,
                    cu = function(t) {
                        var e = Z.Symbol || (Z.Symbol = Te ? {} : b.Symbol || {});
                        "_" == t.charAt(0) || t in e || au(e, t, {
                            value: uu.f(t)
                        })
                    },
                    fu = Object.freeze({
                        default: cu,
                        __moduleExports: cu
                    }),
                    su = function(t) {
                        var e = jn(t),
                            n = xo.f;
                        if (n)
                            for (var r, o = n(t), i = kr.f, u = 0; o.length > u;)
                                i.call(t, r = o[u++]) && e.push(r);
                        return e
                    },
                    lu = Object.freeze({
                        default: su,
                        __moduleExports: su
                    }),
                    hu = Be.f,
                    du = {}.toString,
                    vu = "object" == ("undefined" === typeof window ? "undefined" : a(window)) && window && Object.getOwnPropertyNames ? Object.getOwnPropertyNames(window) : [],
                    pu = function(t) {
                        return vu && "[object Window]" == du.call(t) ? function(t) {
                            try {
                                return hu(t)
                            } catch (t) {
                                return vu.slice()
                            }
                        }(t) : hu(Zt(t))
                    },
                    yu = {
                        f: pu
                    },
                    gu = Object.freeze({
                        default: yu,
                        __moduleExports: yu,
                        f: pu
                    }),
                    mu = fu && cu || fu,
                    bu = lu && su || lu,
                    _u = gu && yu || gu,
                    Ou = eo.KEY,
                    ju = Br.f,
                    Eu = W.f,
                    wu = _u.f,
                    xu = b.Symbol,
                    Su = b.JSON,
                    Mu = Su && Su.stringify,
                    zu = ke("_hidden"),
                    Au = ke("toPrimitive"),
                    Pu = {}.propertyIsEnumerable,
                    ku = ue("symbol-registry"),
                    Lu = ue("symbols"),
                    Fu = ue("op-symbols"),
                    Nu = Object.prototype,
                    Iu = "function" == typeof xu,
                    Tu = b.QObject,
                    Ru = !Tu || !Tu.prototype || !Tu.prototype.findChild,
                    Cu = w && y((function() {
                        return 7 != ar(Eu({}, "a", {
                            get: function() {
                                return Eu(this, "a", {
                                    value: 7
                                }).a
                            }
                        })).a
                    })) ? function(t, e, n) {
                        var r = ju(Nu, e);
                        r && delete Nu[e],
                        Eu(t, e, n),
                        r && t !== Nu && Eu(Nu, e, r)
                    } : Eu,
                    Wu = function(t) {
                        var e = Lu[t] = ar(xu.prototype);
                        return e._k = t, e
                    },
                    Du = Iu && "symbol" == a(xu.iterator) ? function(t) {
                        return "symbol" == a(t)
                    } : function(t) {
                        return t instanceof xu
                    },
                    Bu = function t(e, n, r) {
                        return e === Nu && t(Fu, n, r), P(e), n = L(n, !0), P(r), $(Lu, n) ? (r.enumerable ? ($(e, zu) && e[zu][n] && (e[zu][n] = !1), r = ar(r, {
                            enumerable: D(0, !1)
                        })) : ($(e, zu) || Eu(e, zu, D(1, {})), e[zu][n] = !0), Cu(e, n, r)) : Eu(e, n, r)
                    },
                    Ku = function(t, e) {
                        P(t);
                        for (var n, r = bu(e = Zt(e)), o = 0, i = r.length; i > o;)
                            Bu(t, n = r[o++], e[n]);
                        return t
                    },
                    Vu = function(t) {
                        var e = Pu.call(this, t = L(t, !0));
                        return !(this === Nu && $(Lu, t) && !$(Fu, t)) && (!(e || !$(this, t) || !$(Lu, t) || $(this, zu) && this[zu][t]) || e)
                    },
                    Uu = function(t, e) {
                        if (t = Zt(t), e = L(e, !0), t !== Nu || !$(Lu, e) || $(Fu, e)) {
                            var n = ju(t, e);
                            return !n || !$(Lu, e) || $(t, zu) && t[zu][e] || (n.enumerable = !0), n
                        }
                    },
                    Gu = function(t) {
                        for (var e, n = wu(Zt(t)), r = [], o = 0; n.length > o;)
                            $(Lu, e = n[o++]) || e == zu || e == Ou || r.push(e);
                        return r
                    },
                    Hu = function(t) {
                        for (var e, n = t === Nu, r = wu(n ? Fu : Zt(t)), o = [], i = 0; r.length > i;)
                            !$(Lu, e = r[i++]) || n && !$(Nu, e) || o.push(Lu[e]);
                        return o
                    };
                Iu || (at((xu = function() {
                    if (this instanceof xu)
                        throw TypeError("Symbol is not a constructor!");
                    var t = Q(arguments.length > 0 ? arguments[0] : void 0),
                        e = function e(n) {
                            this === Nu && e.call(Fu, n),
                            $(this, zu) && $(this[zu], t) && (this[zu][t] = !1),
                            Cu(this, t, D(1, n))
                        };
                    return w && Ru && Cu(Nu, t, {
                        configurable: !0,
                        set: e
                    }), Wu(t)
                }).prototype, "toString", (function() {
                    return this._k
                })), Br.f = Uu, W.f = Bu, Be.f = _u.f = Gu, kr.f = Vu, xo.f = Hu, w && !Te && at(Nu, "propertyIsEnumerable", Vu, !0), uu.f = function(t) {
                    return Wu(ke(t))
                }),
                Qe(Qe.G + Qe.W + Qe.F * !Iu, {
                    Symbol: xu
                });
                for (var Ju = "hasInstance,isConcatSpreadable,iterator,match,replace,search,species,split,toPrimitive,toStringTag,unscopables".split(","), Yu = 0; Ju.length > Yu;)
                    ke(Ju[Yu++]);
                for (var qu = jn(ke.store), Xu = 0; qu.length > Xu;)
                    mu(qu[Xu++]);
                Qe(Qe.S + Qe.F * !Iu, "Symbol", {
                    for: function(t) {
                        return $(ku, t += "") ? ku[t] : ku[t] = xu(t)
                    },
                    keyFor: function(t) {
                        if (!Du(t))
                            throw TypeError(t + " is not a symbol!");
                        for (var e in ku)
                            if (ku[e] === t)
                                return e
                    },
                    useSetter: function() {
                        Ru = !0
                    },
                    useSimple: function() {
                        Ru = !1
                    }
                }),
                Qe(Qe.S + Qe.F * !Iu, "Object", {
                    create: function(t, e) {
                        return void 0 === e ? ar(t) : Ku(ar(t), e)
                    },
                    defineProperty: Bu,
                    defineProperties: Ku,
                    getOwnPropertyDescriptor: Uu,
                    getOwnPropertyNames: Gu,
                    getOwnPropertySymbols: Hu
                }),
                Su && Qe(Qe.S + Qe.F * (!Iu || y((function() {
                    var t = xu();
                    return "[null]" != Mu([t]) || "{}" != Mu({
                            a: t
                        }) || "{}" != Mu(Object(t))
                }))), "JSON", {
                    stringify: function(t) {
                        for (var e, n, r = [t], o = 1; arguments.length > o;)
                            r.push(arguments[o++]);
                        if (n = e = r[1], (l(e) || void 0 !== t) && !Du(t))
                            return Gn(e) || (e = function(t, e) {
                                if ("function" == typeof n && (e = n.call(this, t, e)), !Du(e))
                                    return e
                            }), r[1] = e, Mu.apply(Su, r)
                    }
                }),
                xu.prototype[Au] || X(xu.prototype, Au, xu.prototype.valueOf),
                Ve(xu, "Symbol"),
                Ve(Math, "Math", !0),
                Ve(b.JSON, "JSON", !0);
                var $u = function(t, e) {
                        var n = (Z.Object || {})[t] || Object[t],
                            r = {};
                        r[t] = e(n),
                        Qe(Qe.S + Qe.F * y((function() {
                            n(1)
                        })), "Object", r)
                    },
                    Qu = Object.freeze({
                        default: $u,
                        __moduleExports: $u
                    }),
                    Zu = Qu && $u || Qu,
                    ta = eo.onFreeze;
                Zu("freeze", (function(t) {
                    return function(e) {
                        return t && l(e) ? t(ta(e)) : e
                    }
                }));
                var ea = eo.onFreeze;
                Zu("seal", (function(t) {
                    return function(e) {
                        return t && l(e) ? t(ea(e)) : e
                    }
                }));
                var na = eo.onFreeze;
                Zu("preventExtensions", (function(t) {
                    return function(e) {
                        return t && l(e) ? t(na(e)) : e
                    }
                })),
                Zu("isFrozen", (function(t) {
                    return function(e) {
                        return !l(e) || !!t && t(e)
                    }
                })),
                Zu("isSealed", (function(t) {
                    return function(e) {
                        return !l(e) || !!t && t(e)
                    }
                })),
                Zu("isExtensible", (function(t) {
                    return function(e) {
                        return !!l(e) && (!t || t(e))
                    }
                }));
                var ra = Br.f;
                Zu("getOwnPropertyDescriptor", (function() {
                    return function(t, e) {
                        return ra(Zt(t), e)
                    }
                })),
                Zu("getPrototypeOf", (function() {
                    return function(t) {
                        return hr(Se(t))
                    }
                })),
                Zu("keys", (function() {
                    return function(t) {
                        return jn(Se(t))
                    }
                })),
                Zu("getOwnPropertyNames", (function() {
                    return _u.f
                })),
                Qe(Qe.S + Qe.F, "Object", {
                    assign: Co
                });
                var oa = Object.is || function(t, e) {
                        return t === e ? 0 !== t || 1 / t == 1 / e : t != t && e != e
                    },
                    ia = Object.freeze({
                        default: oa,
                        __moduleExports: oa
                    });
                Qe(Qe.S, "Object", {
                    is: ia && oa || ia
                }),
                Qe(Qe.S, "Object", {
                    setPrototypeOf: ho.set
                });
                var ua = Function.prototype,
                    aa = /^\s*function ([^ (]*)/;
                "name" in ua || w && (0, W.f)(ua, "name", {
                    configurable: !0,
                    get: function() {
                        try {
                            return ("" + this).match(aa)[1]
                        } catch (t) {
                            return ""
                        }
                    }
                }),
                Qe(Qe.S, "String", {
                    raw: function(t) {
                        for (var e = Zt(t.raw), n = Nt(e.length), r = arguments.length, o = [], i = 0; n > i;)
                            o.push(e[i++] + ""),
                            r > i && o.push(arguments[i] + "");
                        return o.join("")
                    }
                });
                var ca = String.fromCharCode,
                    fa = String.fromCodePoint;
                Qe(Qe.S + Qe.F * (!!fa && 1 != fa.length), "String", {
                    fromCodePoint: function(t) {
                        for (var e, n = [], r = arguments.length, o = 0; r > o;) {
                            if (e = +arguments[o++], te(e, 1114111) !== e)
                                throw RangeError(e + " is not a valid code point");
                            n.push(65536 > e ? ca(e) : ca(55296 + ((e -= 65536) >> 10), e % 1024 + 56320))
                        }
                        return n.join("")
                    }
                });
                var sa = function(t) {
                        return function(e, n) {
                            var r,
                                o,
                                i = Ht(e) + "",
                                u = Pt(n),
                                a = i.length;
                            return 0 > u || u >= a ? t ? "" : void 0 : 55296 > (r = i.charCodeAt(u)) || r > 56319 || u + 1 === a || 56320 > (o = i.charCodeAt(u + 1)) || o > 57343 ? t ? i.charAt(u) : r : t ? i.slice(u, u + 2) : o - 56320 + (r - 55296 << 10) + 65536
                        }
                    },
                    la = Object.freeze({
                        default: sa,
                        __moduleExports: sa
                    }),
                    ha = (la && sa || la)(!1);
                Qe(Qe.P, "String", {
                    codePointAt: function(t) {
                        return ha(this, t)
                    }
                });
                var da = function(t) {
                        var e = Ht(this) + "",
                            n = "",
                            r = Pt(t);
                        if (0 > r || r == 1 / 0)
                            throw RangeError("Count can't be negative");
                        for (; r > 0; (r >>>= 1) && (e += e))
                            1 & r && (n += e);
                        return n
                    },
                    va = Object.freeze({
                        default: da,
                        __moduleExports: da
                    }),
                    pa = va && da || va;
                Qe(Qe.P, "String", {
                    repeat: pa
                });
                var ya = ke("match"),
                    ga = function(t) {
                        var e;
                        return l(t) && (void 0 !== (e = t[ya]) ? !!e : "RegExp" == Dt(t))
                    },
                    ma = Object.freeze({
                        default: ga,
                        __moduleExports: ga
                    }),
                    ba = ma && ga || ma,
                    _a = function(t, e, n) {
                        if (ba(e))
                            throw TypeError("String#" + n + " doesn't accept regex!");
                        return Ht(t) + ""
                    },
                    Oa = Object.freeze({
                        default: _a,
                        __moduleExports: _a
                    }),
                    ja = ke("match"),
                    Ea = function(t) {
                        var e = /./;
                        try {
                            "/./"[t](e)
                        } catch (r) {
                            try {
                                return e[ja] = !1, !"/./"[t](e)
                            } catch (t) {}
                        }
                        return !0
                    },
                    wa = Object.freeze({
                        default: Ea,
                        __moduleExports: Ea
                    }),
                    xa = Oa && _a || Oa,
                    Sa = wa && Ea || wa,
                    Ma = "".startsWith;
                Qe(Qe.P + Qe.F * Sa("startsWith"), "String", {
                    startsWith: function(t) {
                        var e = xa(this, t, "startsWith"),
                            n = Nt(Math.min(arguments.length > 1 ? arguments[1] : void 0, e.length)),
                            r = t + "";
                        return Ma ? Ma.call(e, r, n) : e.slice(n, n + r.length) === r
                    }
                });
                var za = "".endsWith;
                Qe(Qe.P + Qe.F * Sa("endsWith"), "String", {
                    endsWith: function(t) {
                        var e = xa(this, t, "endsWith"),
                            n = arguments.length > 1 ? arguments[1] : void 0,
                            r = Nt(e.length),
                            o = void 0 === n ? r : Math.min(Nt(n), r),
                            i = t + "";
                        return za ? za.call(e, i, o) : e.slice(o - i.length, o) === i
                    }
                }),
                Qe(Qe.P + Qe.F * Sa("includes"), "String", {
                    includes: function(t) {
                        return !!~xa(this, t, "includes").indexOf(t, arguments.length > 1 ? arguments[1] : void 0)
                    }
                });
                var Aa = function() {
                        var t = P(this),
                            e = "";
                        return t.global && (e += "g"), t.ignoreCase && (e += "i"), t.multiline && (e += "m"), t.unicode && (e += "u"), t.sticky && (e += "y"), e
                    },
                    Pa = Object.freeze({
                        default: Aa,
                        __moduleExports: Aa
                    });
                w && "g" != /./g.flags && W.f(RegExp.prototype, "flags", {
                    configurable: !0,
                    get: Pa && Aa || Pa
                });
                var ka = function(t, e, n) {
                        var r = ke(t),
                            o = n(Ht, r, ""[t]),
                            i = o[0],
                            u = o[1];
                        y((function() {
                            var e = {};
                            return e[r] = function() {
                                return 7
                            }, 7 != ""[t](e)
                        })) && (at(String.prototype, t, i), X(RegExp.prototype, r, 2 == e ? function(t, e) {
                            return u.call(t, this, e)
                        } : function(t) {
                            return u.call(t, this)
                        }))
                    },
                    La = Object.freeze({
                        default: ka,
                        __moduleExports: ka
                    }),
                    Fa = La && ka || La;
                Fa("match", 1, (function(t, e, n) {
                    return [function(n) {
                        var r = t(this),
                            o = void 0 == n ? void 0 : n[e];
                        return void 0 !== o ? o.call(n, r) : RegExp(n)[e](r + "")
                    }, n]
                })),
                Fa("replace", 2, (function(t, e, n) {
                    return [function(r, o) {
                        var i = t(this),
                            u = void 0 == r ? void 0 : r[e];
                        return void 0 !== u ? u.call(r, i, o) : n.call(i + "", r, o)
                    }, n]
                })),
                Fa("split", 2, (function(t, e, n) {
                    var r = ba,
                        o = n,
                        i = [].push;
                    if ("".split(/.?/).length) {
                        var u = void 0 === /()??/.exec("")[1];
                        n = function(t, e) {
                            var n = this + "";
                            if (void 0 === t && 0 === e)
                                return [];
                            if (!r(t))
                                return o.call(n, t, e);
                            var a,
                                c,
                                f,
                                s,
                                l,
                                h = [],
                                d = (t.ignoreCase ? "i" : "") + (t.multiline ? "m" : "") + (t.unicode ? "u" : "") + (t.sticky ? "y" : ""),
                                v = 0,
                                p = void 0 === e ? 4294967295 : e >>> 0,
                                y = RegExp(t.source, d + "g");
                            for (u || (a = RegExp("^" + y.source + "$(?!\\s)", d)); (c = y.exec(n)) && ((f = c.index + c[0].length) <= v || (h.push(n.slice(v, c.index)), !u && c.length > 1 && c[0].replace(a, (function() {
                                for (l = 1; arguments.length - 2 > l; l++)
                                    void 0 === arguments[l] && (c[l] = void 0)
                            })), c.length > 1 && n.length > c.index && i.apply(h, c.slice(1)), s = c[0].length, v = f, p > h.length));)
                                y.lastIndex === c.index && y.lastIndex++;
                            return v === n.length ? !s && y.test("") || h.push("") : h.push(n.slice(v)), h.length > p ? h.slice(0, p) : h
                        }
                    }
                    return [function(r, o) {
                        var i = t(this),
                            u = void 0 == r ? void 0 : r[e];
                        return void 0 !== u ? u.call(r, i, o) : n.call(i + "", r, o)
                    }, n]
                })),
                Fa("search", 1, (function(t, e, n) {
                    return [function(n) {
                        var r = t(this),
                            o = void 0 == n ? void 0 : n[e];
                        return void 0 !== o ? o.call(n, r) : RegExp(n)[e](r + "")
                    }, n]
                }));
                var Na = function(t, e, n) {
                        e in t ? W.f(t, e, D(0, n)) : t[e] = n
                    },
                    Ia = Object.freeze({
                        default: Na,
                        __moduleExports: Na
                    }),
                    Ta = Ia && Na || Ia;
                Qe(Qe.S + Qe.F * !Wr((function(t) {})), "Array", {
                    from: function(t) {
                        var e,
                            n,
                            r,
                            o,
                            i = Se(t),
                            u = "function" == typeof this ? this : Array,
                            a = arguments.length,
                            c = a > 1 ? arguments[1] : void 0,
                            f = void 0 !== c,
                            s = 0,
                            l = Rr(i);
                        if (f && (c = ct(c, a > 2 ? arguments[2] : void 0, 2)), void 0 == l || u == Array && Tr(l))
                            for (n = new u(e = Nt(i.length)); e > s; s++)
                                Ta(n, s, f ? c(i[s], s) : i[s]);
                        else
                            for (o = l.call(i), n = new u; !(r = o.next()).done; s++)
                                Ta(n, s, f ? Jr(o, c, [r.value, s], !0) : r.value);
                        return n.length = s, n
                    }
                }),
                Qe(Qe.S + Qe.F * y((function() {
                    function t() {}
                    return !(Array.of.call(t) instanceof t)
                })), "Array", {
                    of: function() {
                        for (var t = 0, e = arguments.length, n = new ("function" == typeof this ? this : Array)(e); e > t;)
                            Ta(n, t, arguments[t++]);
                        return n.length = e, n
                    }
                }),
                Qe(Qe.P, "Array", {
                    copyWithin: Dr
                }),
                mr("copyWithin");
                var Ra = Cr(5),
                    Ca = !0;
                "find" in [] && Array(1).find((function() {
                    Ca = !1
                })),
                Qe(Qe.P + Qe.F * Ca, "Array", {
                    find: function(t) {
                        return Ra(this, t, arguments.length > 1 ? arguments[1] : void 0)
                    }
                }),
                mr("find");
                var Wa = Cr(6),
                    Da = !0;
                "findIndex" in [] && Array(1).findIndex((function() {
                    Da = !1
                })),
                Qe(Qe.P + Qe.F * Da, "Array", {
                    findIndex: function(t) {
                        return Wa(this, t, arguments.length > 1 ? arguments[1] : void 0)
                    }
                }),
                mr("findIndex"),
                Qe(Qe.P, "Array", {
                    fill: Ke
                }),
                mr("fill");
                var Ba = b.isFinite;
                Qe(Qe.S, "Number", {
                    isFinite: function(t) {
                        return "number" == typeof t && Ba(t)
                    }
                });
                var Ka = Math.floor,
                    Va = function(t) {
                        return !l(t) && isFinite(t) && Ka(t) === t
                    },
                    Ua = Object.freeze({
                        default: Va,
                        __moduleExports: Va
                    }),
                    Ga = Ua && Va || Ua;
                Qe(Qe.S, "Number", {
                    isInteger: Ga
                });
                var Ha = Math.abs;
                Qe(Qe.S, "Number", {
                    isSafeInteger: function(t) {
                        return Ga(t) && 9007199254740991 >= Ha(t)
                    }
                }),
                Qe(Qe.S, "Number", {
                    isNaN: function(t) {
                        return t != t
                    }
                }),
                Qe(Qe.S, "Number", {
                    EPSILON: Math.pow(2, -52)
                }),
                Qe(Qe.S, "Number", {
                    MIN_SAFE_INTEGER: -9007199254740991
                }),
                Qe(Qe.S, "Number", {
                    MAX_SAFE_INTEGER: 9007199254740991
                });
                var Ja = Math.log1p || function(t) {
                        return (t = +t) > -1e-8 && 1e-8 > t ? t - t * t / 2 : Math.log(1 + t)
                    },
                    Ya = Object.freeze({
                        default: Ja,
                        __moduleExports: Ja
                    }),
                    qa = Ya && Ja || Ya,
                    Xa = Math.sqrt,
                    $a = Math.acosh;
                Qe(Qe.S + Qe.F * !($a && 710 == Math.floor($a(Number.MAX_VALUE)) && $a(1 / 0) == 1 / 0), "Math", {
                    acosh: function(t) {
                        return 1 > (t = +t) ? NaN : t > 94906265.62425156 ? Math.log(t) + Math.LN2 : qa(t - 1 + Xa(t - 1) * Xa(t + 1))
                    }
                });
                var Qa = Math.asinh;
                Qe(Qe.S + Qe.F * !(Qa && 1 / Qa(0) > 0), "Math", {
                    asinh: function t(e) {
                        return isFinite(e = +e) && 0 != e ? 0 > e ? -t(-e) : Math.log(e + Math.sqrt(e * e + 1)) : e
                    }
                });
                var Za = Math.atanh;
                Qe(Qe.S + Qe.F * !(Za && 0 > 1 / Za(-0)), "Math", {
                    atanh: function(t) {
                        return 0 == (t = +t) ? t : Math.log((1 + t) / (1 - t)) / 2
                    }
                });
                var tc = Math.sign || function(t) {
                        return 0 == (t = +t) || t != t ? t : 0 > t ? -1 : 1
                    },
                    ec = Object.freeze({
                        default: tc,
                        __moduleExports: tc
                    }),
                    nc = ec && tc || ec;
                Qe(Qe.S, "Math", {
                    cbrt: function(t) {
                        return nc(t = +t) * Math.pow(Math.abs(t), 1 / 3)
                    }
                }),
                Qe(Qe.S, "Math", {
                    clz32: function(t) {
                        return (t >>>= 0) ? 31 - Math.floor(Math.log(t + .5) * Math.LOG2E) : 32
                    }
                });
                var rc = Math.exp;
                Qe(Qe.S, "Math", {
                    cosh: function(t) {
                        return (rc(t = +t) + rc(-t)) / 2
                    }
                });
                var oc = Math.expm1,
                    ic = !oc || oc(10) > 22025.465794806718 || 22025.465794806718 > oc(10) || -2e-17 != oc(-2e-17) ? function(t) {
                        return 0 == (t = +t) ? t : t > -1e-6 && 1e-6 > t ? t + t * t / 2 : Math.exp(t) - 1
                    } : oc,
                    uc = Object.freeze({
                        default: ic,
                        __moduleExports: ic
                    }),
                    ac = uc && ic || uc;
                Qe(Qe.S + Qe.F * (ac != Math.expm1), "Math", {
                    expm1: ac
                });
                var cc = Math.pow,
                    fc = cc(2, -52),
                    sc = cc(2, -23),
                    lc = cc(2, 127) * (2 - sc),
                    hc = cc(2, -126),
                    dc = Math.fround || function(t) {
                        var e,
                            n,
                            r = Math.abs(t),
                            o = nc(t);
                        return hc > r ? o * (r / hc / sc + 1 / fc - 1 / fc) * hc * sc : (n = (e = (1 + sc / fc) * r) - (e - r)) > lc || n != n ? o * (1 / 0) : o * n
                    },
                    vc = Object.freeze({
                        default: dc,
                        __moduleExports: dc
                    });
                Qe(Qe.S, "Math", {
                    fround: vc && dc || vc
                });
                var pc = Math.abs;
                Qe(Qe.S, "Math", {
                    hypot: function(t, e) {
                        for (var n, r, o = 0, i = 0, u = arguments.length, a = 0; u > i;)
                            (n = pc(arguments[i++])) > a ? (o = o * (r = a / n) * r + 1, a = n) : o += n > 0 ? (r = n / a) * r : n;
                        return a === 1 / 0 ? 1 / 0 : a * Math.sqrt(o)
                    }
                });
                var yc = Math.imul;
                Qe(Qe.S + Qe.F * y((function() {
                    return -5 != yc(4294967295, 5) || 2 != yc.length
                })), "Math", {
                    imul: function(t, e) {
                        var n = +t,
                            r = +e,
                            o = 65535 & n,
                            i = 65535 & r;
                        return 0 | o * i + ((65535 & n >>> 16) * i + o * (65535 & r >>> 16) << 16 >>> 0)
                    }
                }),
                Qe(Qe.S, "Math", {
                    log1p: qa
                }),
                Qe(Qe.S, "Math", {
                    log10: function(t) {
                        return Math.log(t) * Math.LOG10E
                    }
                }),
                Qe(Qe.S, "Math", {
                    log2: function(t) {
                        return Math.log(t) / Math.LN2
                    }
                }),
                Qe(Qe.S, "Math", {
                    sign: nc
                });
                var gc = Math.exp;
                Qe(Qe.S + Qe.F * y((function() {
                    return -2e-17 != !Math.sinh(-2e-17)
                })), "Math", {
                    sinh: function(t) {
                        return 1 > Math.abs(t = +t) ? (ac(t) - ac(-t)) / 2 : (gc(t - 1) - gc(-t - 1)) * (Math.E / 2)
                    }
                });
                var mc = Math.exp;
                Qe(Qe.S, "Math", {
                    tanh: function(t) {
                        var e = ac(t = +t),
                            n = ac(-t);
                        return e == 1 / 0 ? 1 : n == 1 / 0 ? -1 : (e - n) / (mc(t) + mc(-t))
                    }
                }),
                Qe(Qe.S, "Math", {
                    trunc: function(t) {
                        return (t > 0 ? Math.floor : Math.ceil)(t)
                    }
                });
                var bc = se(!0);
                Qe(Qe.P, "Array", {
                    includes: function(t) {
                        return bc(this, t, arguments.length > 1 ? arguments[1] : void 0)
                    }
                }),
                mr("includes");
                var _c = kr.f,
                    Oc = function(t) {
                        return function(e) {
                            for (var n, r = Zt(e), o = jn(r), i = o.length, u = 0, a = []; i > u;)
                                _c.call(r, n = o[u++]) && a.push(t ? [n, r[n]] : r[n]);
                            return a
                        }
                    },
                    jc = Object.freeze({
                        default: Oc,
                        __moduleExports: Oc
                    }),
                    Ec = jc && Oc || jc,
                    wc = Ec(!1);
                Qe(Qe.S, "Object", {
                    values: function(t) {
                        return wc(t)
                    }
                });
                var xc = Ec(!0);
                Qe(Qe.S, "Object", {
                    entries: function(t) {
                        return xc(t)
                    }
                }),
                Qe(Qe.S, "Object", {
                    getOwnPropertyDescriptors: function(t) {
                        for (var e, n, r = Zt(t), o = Br.f, i = oi(r), u = {}, a = 0; i.length > a;)
                            void 0 !== (n = o(r, e = i[a++])) && Ta(u, e, n);
                        return u
                    }
                });
                var Sc = function(t, e, n, r) {
                        var o = Ht(t) + "",
                            i = o.length,
                            u = void 0 === n ? " " : n + "",
                            a = Nt(e);
                        if (i >= a || "" == u)
                            return o;
                        var c = a - i,
                            f = pa.call(u, Math.ceil(c / u.length));
                        return f.length > c && (f = f.slice(0, c)), r ? f + o : o + f
                    },
                    Mc = Object.freeze({
                        default: Sc,
                        __moduleExports: Sc
                    }),
                    zc = b.navigator,
                    Ac = zc && zc.userAgent || "",
                    Pc = Object.freeze({
                        default: Ac,
                        __moduleExports: Ac
                    }),
                    kc = Mc && Sc || Mc,
                    Lc = Pc && Ac || Pc;
                Qe(Qe.P + Qe.F * /Version\/10\.\d+(\.\d+)? Safari\//.test(Lc), "String", {
                    padStart: function(t) {
                        return kc(this, t, arguments.length > 1 ? arguments[1] : void 0, !0)
                    }
                }),
                Qe(Qe.P + Qe.F * /Version\/10\.\d+(\.\d+)? Safari\//.test(Lc), "String", {
                    padEnd: function(t) {
                        return kc(this, t, arguments.length > 1 ? arguments[1] : void 0, !1)
                    }
                });
                var Fc = [].slice,
                    Nc = function(t) {
                        return function(e, n) {
                            var r = arguments.length > 2,
                                o = !!r && Fc.call(arguments, 2);
                            return t(r ? function() {
                                ("function" == typeof e ? e : Function(e)).apply(this, o)
                            } : e, n)
                        }
                    };
                Qe(Qe.G + Qe.B + Qe.F * /MSIE .\./.test(Lc), {
                    setTimeout: Nc(b.setTimeout),
                    setInterval: Nc(b.setInterval)
                }),
                Qe(Qe.G + Qe.B, {
                    setImmediate: wi.set,
                    clearImmediate: wi.clear
                });
                for (var Ic = ke("iterator"), Tc = ke("toStringTag"), Rc = pn.Array, Cc = {
                        CSSRuleList: !0,
                        CSSStyleDeclaration: !1,
                        CSSValueList: !1,
                        ClientRectList: !1,
                        DOMRectList: !1,
                        DOMStringList: !1,
                        DOMTokenList: !0,
                        DataTransferItemList: !1,
                        FileList: !1,
                        HTMLAllCollection: !1,
                        HTMLCollection: !1,
                        HTMLFormElement: !1,
                        HTMLSelectElement: !1,
                        MediaList: !0,
                        MimeTypeArray: !1,
                        NamedNodeMap: !1,
                        NodeList: !0,
                        PaintRequestList: !1,
                        Plugin: !1,
                        PluginArray: !1,
                        SVGLengthList: !1,
                        SVGNumberList: !1,
                        SVGPathSegList: !1,
                        SVGPointList: !1,
                        SVGStringList: !1,
                        SVGTransformList: !1,
                        SourceBufferList: !1,
                        StyleSheetList: !0,
                        TextTrackCueList: !1,
                        TextTrackList: !1,
                        TouchList: !1
                    }, Wc = jn(Cc), Dc = 0; Wc.length > Dc; Dc++) {
                    var Bc,
                        Kc = Wc[Dc],
                        Vc = Cc[Kc],
                        Uc = b[Kc],
                        Gc = Uc && Uc.prototype;
                    if (Gc && (Gc[Ic] || X(Gc, Ic, Rc), Gc[Tc] || X(Gc, Tc, Kc), pn[Kc] = Rc, Vc))
                        for (Bc in Or)
                            Gc[Bc] || at(Gc, Bc, Or[Bc], !0)
                }
                r((function(t) {
                    !function(e) {
                        var n,
                            r = Object.prototype,
                            o = r.hasOwnProperty,
                            i = "function" == typeof Symbol ? Symbol : {},
                            u = i.iterator || "@@iterator",
                            c = i.asyncIterator || "@@asyncIterator",
                            f = i.toStringTag || "@@toStringTag",
                            s = e.regeneratorRuntime;
                        if (s)
                            t.exports = s;
                        else {
                            (s = e.regeneratorRuntime = t.exports).wrap = _;
                            var l = "suspendedStart",
                                h = "suspendedYield",
                                d = "executing",
                                v = "completed",
                                p = {},
                                y = {};
                            y[u] = function() {
                                return this
                            };
                            var g = Object.getPrototypeOf,
                                m = g && g(g(k([])));
                            m && m !== r && o.call(m, u) && (y = m);
                            var b = w.prototype = j.prototype = Object.create(y);
                            E.prototype = b.constructor = w,
                            w.constructor = E,
                            w[f] = E.displayName = "GeneratorFunction",
                            s.isGeneratorFunction = function(t) {
                                var e = "function" == typeof t && t.constructor;
                                return !!e && (e === E || "GeneratorFunction" === (e.displayName || e.name))
                            },
                            s.mark = function(t) {
                                return Object.setPrototypeOf ? Object.setPrototypeOf(t, w) : (t.__proto__ = w, f in t || (t[f] = "GeneratorFunction")), t.prototype = Object.create(b), t
                            },
                            s.awrap = function(t) {
                                return {
                                    __await: t
                                }
                            },
                            x(S.prototype),
                            S.prototype[c] = function() {
                                return this
                            },
                            s.AsyncIterator = S,
                            s.async = function(t, e, n, r) {
                                var o = new S(_(t, e, n, r));
                                return s.isGeneratorFunction(e) ? o : o.next().then((function(t) {
                                    return t.done ? t.value : o.next()
                                }))
                            },
                            x(b),
                            b[f] = "Generator",
                            b[u] = function() {
                                return this
                            },
                            b.toString = function() {
                                return "[object Generator]"
                            },
                            s.keys = function(t) {
                                var e = [];
                                for (var n in t)
                                    e.push(n);
                                return e.reverse(), function n() {
                                    for (; e.length;) {
                                        var r = e.pop();
                                        if (r in t)
                                            return n.value = r, n.done = !1, n
                                    }
                                    return n.done = !0, n
                                }
                            },
                            s.values = k,
                            P.prototype = {
                                constructor: P,
                                reset: function(t) {
                                    if (this.prev = 0, this.next = 0, this.sent = this._sent = n, this.done = !1, this.delegate = null, this.method = "next", this.arg = n, this.tryEntries.forEach(A), !t)
                                        for (var e in this)
                                            "t" === e.charAt(0) && o.call(this, e) && !isNaN(+e.slice(1)) && (this[e] = n)
                                },
                                stop: function() {
                                    this.done = !0;
                                    var t = this.tryEntries[0].completion;
                                    if ("throw" === t.type)
                                        throw t.arg;
                                    return this.rval
                                },
                                dispatchException: function(t) {
                                    if (this.done)
                                        throw t;
                                    var e = this;
                                    function r(r, o) {
                                        return a.type = "throw", a.arg = t, e.next = r, o && (e.method = "next", e.arg = n), !!o
                                    }
                                    for (var i = this.tryEntries.length - 1; i >= 0; --i) {
                                        var u = this.tryEntries[i],
                                            a = u.completion;
                                        if ("root" === u.tryLoc)
                                            return r("end");
                                        if (this.prev >= u.tryLoc) {
                                            var c = o.call(u, "catchLoc"),
                                                f = o.call(u, "finallyLoc");
                                            if (c && f) {
                                                if (u.catchLoc > this.prev)
                                                    return r(u.catchLoc, !0);
                                                if (u.finallyLoc > this.prev)
                                                    return r(u.finallyLoc)
                                            } else if (c) {
                                                if (u.catchLoc > this.prev)
                                                    return r(u.catchLoc, !0)
                                            } else {
                                                if (!f)
                                                    throw Error("try statement without catch or finally");
                                                if (u.finallyLoc > this.prev)
                                                    return r(u.finallyLoc)
                                            }
                                        }
                                    }
                                },
                                abrupt: function(t, e) {
                                    for (var n = this.tryEntries.length - 1; n >= 0; --n) {
                                        var r = this.tryEntries[n];
                                        if (this.prev >= r.tryLoc && o.call(r, "finallyLoc") && r.finallyLoc > this.prev) {
                                            var i = r;
                                            break
                                        }
                                    }
                                    !i || "break" !== t && "continue" !== t || i.tryLoc > e || e > i.finallyLoc || (i = null);
                                    var u = i ? i.completion : {};
                                    return u.type = t, u.arg = e, i ? (this.method = "next", this.next = i.finallyLoc, p) : this.complete(u)
                                },
                                complete: function(t, e) {
                                    if ("throw" === t.type)
                                        throw t.arg;
                                    return "break" === t.type || "continue" === t.type ? this.next = t.arg : "return" === t.type ? (this.rval = this.arg = t.arg, this.method = "return", this.next = "end") : "normal" === t.type && e && (this.next = e), p
                                },
                                finish: function(t) {
                                    for (var e = this.tryEntries.length - 1; e >= 0; --e) {
                                        var n = this.tryEntries[e];
                                        if (n.finallyLoc === t)
                                            return this.complete(n.completion, n.afterLoc), A(n), p
                                    }
                                },
                                catch: function(t) {
                                    for (var e = this.tryEntries.length - 1; e >= 0; --e) {
                                        var n = this.tryEntries[e];
                                        if (n.tryLoc === t) {
                                            var r = n.completion;
                                            if ("throw" === r.type) {
                                                var o = r.arg;
                                                A(n)
                                            }
                                            return o
                                        }
                                    }
                                    throw Error("illegal catch attempt")
                                },
                                delegateYield: function(t, e, r) {
                                    return this.delegate = {
                                        iterator: k(t),
                                        resultName: e,
                                        nextLoc: r
                                    }, "next" === this.method && (this.arg = n), p
                                }
                            }
                        }
                        function _(t, e, n, r) {
                            var o = Object.create((e && e.prototype instanceof j ? e : j).prototype),
                                i = new P(r || []);
                            return o._invoke = function(t, e, n) {
                                var r = l;
                                return function(o, i) {
                                    if (r === d)
                                        throw Error("Generator is already running");
                                    if (r === v) {
                                        if ("throw" === o)
                                            throw i;
                                        return L()
                                    }
                                    for (n.method = o, n.arg = i;;) {
                                        var u = n.delegate;
                                        if (u) {
                                            var a = M(u, n);
                                            if (a) {
                                                if (a === p)
                                                    continue;
                                                return a
                                            }
                                        }
                                        if ("next" === n.method)
                                            n.sent = n._sent = n.arg;
                                        else if ("throw" === n.method) {
                                            if (r === l)
                                                throw r = v, n.arg;
                                            n.dispatchException(n.arg)
                                        } else
                                            "return" === n.method && n.abrupt("return", n.arg);
                                        r = d;
                                        var c = O(t, e, n);
                                        if ("normal" === c.type) {
                                            if (r = n.done ? v : h, c.arg === p)
                                                continue;
                                            return {
                                                value: c.arg,
                                                done: n.done
                                            }
                                        }
                                        "throw" === c.type && (r = v, n.method = "throw", n.arg = c.arg)
                                    }
                                }
                            }(t, n, i), o
                        }
                        function O(t, e, n) {
                            try {
                                return {
                                    type: "normal",
                                    arg: t.call(e, n)
                                }
                            } catch (t) {
                                return {
                                    type: "throw",
                                    arg: t
                                }
                            }
                        }
                        function j() {}
                        function E() {}
                        function w() {}
                        function x(t) {
                            ["next", "throw", "return"].forEach((function(e) {
                                t[e] = function(t) {
                                    return this._invoke(e, t)
                                }
                            }))
                        }
                        function S(t) {
                            function n(e, r, i, u) {
                                var c = O(t[e], t, r);
                                if ("throw" !== c.type) {
                                    var f = c.arg,
                                        s = f.value;
                                    return s && "object" == a(s) && o.call(s, "__await") ? Promise.resolve(s.__await).then((function(t) {
                                        n("next", t, i, u)
                                    }), (function(t) {
                                        n("throw", t, i, u)
                                    })) : Promise.resolve(s).then((function(t) {
                                        f.value = t,
                                        i(f)
                                    }), u)
                                }
                                u(c.arg)
                            }
                            var r;
                            "object" == a(e.process) && e.process.domain && (n = e.process.domain.bind(n)),
                            this._invoke = function(t, e) {
                                function o() {
                                    return new Promise((function(r, o) {
                                        n(t, e, r, o)
                                    }))
                                }
                                return r = r ? r.then(o, o) : o()
                            }
                        }
                        function M(t, e) {
                            var r = t.iterator[e.method];
                            if (r === n) {
                                if (e.delegate = null, "throw" === e.method) {
                                    if (t.iterator.return && (e.method = "return", e.arg = n, M(t, e), "throw" === e.method))
                                        return p;
                                    e.method = "throw",
                                    e.arg = new TypeError("The iterator does not provide a 'throw' method")
                                }
                                return p
                            }
                            var o = O(r, t.iterator, e.arg);
                            if ("throw" === o.type)
                                return e.method = "throw", e.arg = o.arg, e.delegate = null, p;
                            var i = o.arg;
                            return i ? i.done ? (e[t.resultName] = i.value, e.next = t.nextLoc, "return" !== e.method && (e.method = "next", e.arg = n), e.delegate = null, p) : i : (e.method = "throw", e.arg = new TypeError("iterator result is not an object"), e.delegate = null, p)
                        }
                        function z(t) {
                            var e = {
                                tryLoc: t[0]
                            };
                            1 in t && (e.catchLoc = t[1]),
                            2 in t && (e.finallyLoc = t[2], e.afterLoc = t[3]),
                            this.tryEntries.push(e)
                        }
                        function A(t) {
                            var e = t.completion || {};
                            e.type = "normal",
                            delete e.arg,
                            t.completion = e
                        }
                        function P(t) {
                            this.tryEntries = [{
                                tryLoc: "root"
                            }],
                            t.forEach(z, this),
                            this.reset(!0)
                        }
                        function k(t) {
                            if (t) {
                                var e = t[u];
                                if (e)
                                    return e.call(t);
                                if ("function" == typeof t.next)
                                    return t;
                                if (!isNaN(t.length)) {
                                    var r = -1,
                                        i = function e() {
                                            for (; ++r < t.length;)
                                                if (o.call(t, r))
                                                    return e.value = t[r], e.done = !1, e;
                                            return e.value = n, e.done = !0, e
                                        };
                                    return i.next = i
                                }
                            }
                            return {
                                next: L
                            }
                        }
                        function L() {
                            return {
                                value: n,
                                done: !0
                            }
                        }
                    }("object" == a(e) ? e : "object" == ("undefined" === typeof window ? "undefined" : a(window)) ? window : "object" == ("undefined" === typeof self ? "undefined" : a(self)) ? self : e)
                }));
                var Hc = "function" == typeof Symbol && "symbol" == a(Symbol.iterator) ? function(t) {
                        return a(t)
                    } : function(t) {
                        return t && "function" == typeof Symbol && t.constructor === Symbol && t !== Symbol.prototype ? "symbol" : a(t)
                    },
                    Jc = function(t, e) {
                        if (Array.isArray(t))
                            return t;
                        if (Symbol.iterator in Object(t))
                            return function(t, e) {
                                var n = [],
                                    r = !0,
                                    o = !1,
                                    i = void 0;
                                try {
                                    for (var u, a = t[Symbol.iterator](); !(r = (u = a.next()).done) && (n.push(u.value), !e || n.length !== e); r = !0)
                                        ;
                                } catch (t) {
                                    o = !0,
                                    i = t
                                } finally {
                                    try {
                                        !r && a.return && a.return()
                                    } finally {
                                        if (o)
                                            throw i
                                    }
                                }
                                return n
                            }(t, e);
                        throw new TypeError("Invalid attempt to destructure non-iterable instance")
                    },
                    Yc = function(t) {
                        if (Array.isArray(t)) {
                            for (var e = 0, n = Array(t.length); t.length > e; e++)
                                n[e] = t[e];
                            return n
                        }
                        return Array.from(t)
                    };
                function qc(t) {
                    return null === t ? "null" : t !== Object(t) ? void 0 === t ? "undefined" : Hc(t) : {}.toString.call(t).slice(8, -1).toLowerCase()
                }
                function Xc(t) {
                    return "string" !== qc(t) || !t.length
                }
                function $c() {
                    var t = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : "",
                        e = arguments[1],
                        n = arguments[2];
                    if (Xc(t))
                        return !1;
                    var r = t.charCodeAt(0);
                    return r >= e && n >= r
                }
                var Qc = {
                        HIRAGANA: "toHiragana",
                        KATAKANA: "toKatakana"
                    },
                    Zc = {
                        HEPBURN: "hepburn"
                    },
                    tf = {
                        useObsoleteKana: !1,
                        passRomaji: !1,
                        upcaseKatakana: !1,
                        ignoreCase: !1,
                        IMEMode: !1,
                        romanization: Zc.HEPBURN
                    },
                    ef = 12353,
                    nf = 12449,
                    rf = [65377, 65381],
                    of = [[12288, 12351], rf, [12539, 12540], [65281, 65295], [65306, 65311], [65339, 65343], [65371, 65376], [65504, 65518]],
                    uf = [].concat([[12352, 12447], [12448, 12543], rf, [65382, 65439]], of, [[65313, 65338], [65345, 65370], [65296, 65305], [19968, 40959], [13312, 19903]]),
                    af = [[0, 127]].concat([[256, 257], [274, 275], [298, 299], [332, 333], [362, 363]]),
                    cf = [[32, 47], [58, 63], [91, 96], [123, 126]].concat([[8216, 8217], [8220, 8221]]);
                function ff() {
                    var t = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : "";
                    return uf.some((function(e) {
                        var n = Jc(e, 2);
                        return $c(t, n[0], n[1])
                    }))
                }
                function sf() {
                    var t = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : "",
                        e = arguments[1],
                        n = "regexp" === qc(e);
                    return !Xc(t) && [].concat(Yc(t)).every((function(t) {
                            var r = ff(t);
                            return n ? r || e.test(t) : r
                        }))
                }
                var lf = function() {
                    return Object.assign({}, tf, arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : {})
                };
                function hf(t, e, n) {
                    var r = e;
                    return function t(e, o) {
                        var i = e.charAt(0);
                        return function e(r, o, i, u) {
                            if (!o)
                                return n || 1 === Object.keys(r).length ? r[""] ? [[i, u, r[""]]] : [] : [[i, u, null]];
                            if (1 === Object.keys(r).length)
                                return [[i, u, r[""]]].concat(t(o, u));
                            var a = function(t, e) {
                                if (void 0 !== t[e])
                                    return Object.assign({
                                        "": t[""] + e
                                    }, t[e])
                            }(r, o.charAt(0));
                            return void 0 === a ? [[i, u, r[""]]].concat(t(o, u)) : e(a, o.slice(1), i, u + 1)
                        }(Object.assign({
                            "": i
                        }, r[i]), e.slice(1), o, o + 1)
                    }(t, 0)
                }
                function df(t) {
                    return Object.entries(t).reduce((function(t, e) {
                        var n = Jc(e, 2),
                            r = n[0],
                            o = n[1],
                            i = "string" === qc(o);
                        return t[r] = i ? {
                            "": o
                        } : df(o), t
                    }), {})
                }
                function vf(t, e) {
                    return e.split("").reduce((function(t, e) {
                        return void 0 === t[e] && (t[e] = {}), t[e]
                    }), t)
                }
                function pf() {
                    var t = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : {},
                        e = {};
                    return "object" === qc(t) && Object.entries(t).forEach((function(t) {
                        var n = Jc(t, 2),
                            r = n[1],
                            o = e;
                        n[0].split("").forEach((function(t) {
                            void 0 === o[t] && (o[t] = {}),
                            o = o[t]
                        })),
                        o[""] = r
                    })), function(t) {
                        return function t(e, n) {
                            return void 0 === e || "string" === qc(e) ? n : Object.entries(n).reduce((function(n, r) {
                                var o = Jc(r, 2),
                                    i = o[0];
                                return n[i] = t(e[i], o[1]), n
                            }), e)
                        }(JSON.parse(JSON.stringify(t)), e)
                    }
                }
                function yf(t, e) {
                    return e ? "function" === qc(e) ? e(t) : pf(e)(t) : t
                }
                var gf = {
                        a: "\u3042",
                        i: "\u3044",
                        u: "\u3046",
                        e: "\u3048",
                        o: "\u304a",
                        k: {
                            a: "\u304b",
                            i: "\u304d",
                            u: "\u304f",
                            e: "\u3051",
                            o: "\u3053"
                        },
                        s: {
                            a: "\u3055",
                            i: "\u3057",
                            u: "\u3059",
                            e: "\u305b",
                            o: "\u305d"
                        },
                        t: {
                            a: "\u305f",
                            i: "\u3061",
                            u: "\u3064",
                            e: "\u3066",
                            o: "\u3068"
                        },
                        n: {
                            a: "\u306a",
                            i: "\u306b",
                            u: "\u306c",
                            e: "\u306d",
                            o: "\u306e"
                        },
                        h: {
                            a: "\u306f",
                            i: "\u3072",
                            u: "\u3075",
                            e: "\u3078",
                            o: "\u307b"
                        },
                        m: {
                            a: "\u307e",
                            i: "\u307f",
                            u: "\u3080",
                            e: "\u3081",
                            o: "\u3082"
                        },
                        y: {
                            a: "\u3084",
                            u: "\u3086",
                            o: "\u3088"
                        },
                        r: {
                            a: "\u3089",
                            i: "\u308a",
                            u: "\u308b",
                            e: "\u308c",
                            o: "\u308d"
                        },
                        w: {
                            a: "\u308f",
                            i: "\u3090",
                            e: "\u3091",
                            o: "\u3092"
                        },
                        g: {
                            a: "\u304c",
                            i: "\u304e",
                            u: "\u3050",
                            e: "\u3052",
                            o: "\u3054"
                        },
                        z: {
                            a: "\u3056",
                            i: "\u3058",
                            u: "\u305a",
                            e: "\u305c",
                            o: "\u305e"
                        },
                        d: {
                            a: "\u3060",
                            i: "\u3062",
                            u: "\u3065",
                            e: "\u3067",
                            o: "\u3069"
                        },
                        b: {
                            a: "\u3070",
                            i: "\u3073",
                            u: "\u3076",
                            e: "\u3079",
                            o: "\u307c"
                        },
                        p: {
                            a: "\u3071",
                            i: "\u3074",
                            u: "\u3077",
                            e: "\u307a",
                            o: "\u307d"
                        },
                        v: {
                            a: "\u3094\u3041",
                            i: "\u3094\u3043",
                            u: "\u3094",
                            e: "\u3094\u3047",
                            o: "\u3094\u3049"
                        }
                    },
                    mf = {
                        ".": "\u3002",
                        ",": "\u3001",
                        ":": "\uff1a",
                        "/": "\u30fb",
                        "!": "\uff01",
                        "?": "\uff1f",
                        "~": "\u301c",
                        "-": "\u30fc",
                        "\u2018": "\u300c",
                        "\u2019": "\u300d",
                        "\u201c": "\u300e",
                        "\u201d": "\u300f",
                        "[": "\uff3b",
                        "]": "\uff3d",
                        "(": "\uff08",
                        ")": "\uff09",
                        "{": "\uff5b",
                        "}": "\uff5d"
                    },
                    bf = {
                        k: "\u304d",
                        s: "\u3057",
                        t: "\u3061",
                        n: "\u306b",
                        h: "\u3072",
                        m: "\u307f",
                        r: "\u308a",
                        g: "\u304e",
                        z: "\u3058",
                        d: "\u3062",
                        b: "\u3073",
                        p: "\u3074",
                        v: "\u3094",
                        q: "\u304f",
                        f: "\u3075"
                    },
                    _f = {
                        ya: "\u3083",
                        yi: "\u3043",
                        yu: "\u3085",
                        ye: "\u3047",
                        yo: "\u3087"
                    },
                    Of = {
                        a: "\u3041",
                        i: "\u3043",
                        u: "\u3045",
                        e: "\u3047",
                        o: "\u3049"
                    },
                    jf = {
                        sh: "sy",
                        ch: "ty",
                        cy: "ty",
                        chy: "ty",
                        shy: "sy",
                        j: "zy",
                        jy: "zy",
                        shi: "si",
                        chi: "ti",
                        tsu: "tu",
                        ji: "zi",
                        fu: "hu"
                    },
                    Ef = Object.assign({
                        tu: "\u3063",
                        wa: "\u308e",
                        ka: "\u30f5",
                        ke: "\u30f6"
                    }, Of, _f),
                    wf = {
                        yi: "\u3044",
                        wu: "\u3046",
                        ye: "\u3044\u3047",
                        wi: "\u3046\u3043",
                        we: "\u3046\u3047",
                        kwa: "\u304f\u3041",
                        whu: "\u3046",
                        tha: "\u3066\u3083",
                        thu: "\u3066\u3085",
                        tho: "\u3066\u3087",
                        dha: "\u3067\u3083",
                        dhu: "\u3067\u3085",
                        dho: "\u3067\u3087"
                    },
                    xf = {
                        wh: "\u3046",
                        qw: "\u304f",
                        q: "\u304f",
                        gw: "\u3050",
                        sw: "\u3059",
                        ts: "\u3064",
                        th: "\u3066",
                        tw: "\u3068",
                        dh: "\u3067",
                        dw: "\u3069",
                        fw: "\u3075",
                        f: "\u3075"
                    };
                function Sf() {
                    var t = df(gf),
                        e = function(e) {
                            return vf(t, e)
                        };
                    return Object.entries(bf).forEach((function(t) {
                        var n = Jc(t, 2),
                            r = n[0],
                            o = n[1];
                        Object.entries(_f).forEach((function(t) {
                            var n = Jc(t, 2),
                                i = n[1];
                            e(r + n[0])[""] = o + i
                        }))
                    })), Object.entries(mf).forEach((function(t) {
                        var n = Jc(t, 2),
                            r = n[1];
                        e(n[0])[""] = r
                    })), Object.entries(xf).forEach((function(t) {
                        var n = Jc(t, 2),
                            r = n[0],
                            o = n[1];
                        Object.entries(Of).forEach((function(t) {
                            var n = Jc(t, 2),
                                i = n[1];
                            e(r + n[0])[""] = o + i
                        }))
                    })), ["n", "n'", "xn"].forEach((function(t) {
                        e(t)[""] = "\u3093"
                    })), t.c = JSON.parse(JSON.stringify(t.k)), Object.entries(jf).forEach((function(t) {
                        var n = Jc(t, 2),
                            r = n[0],
                            o = n[1],
                            i = r.slice(0, r.length - 1),
                            u = r.charAt(r.length - 1);
                        e(i)[u] = JSON.parse(JSON.stringify(e(o)))
                    })), Object.entries(Ef).forEach((function(t) {
                        var n,
                            r = Jc(t, 2),
                            o = r[0],
                            i = r[1],
                            u = function(t) {
                                return t.charAt(t.length - 1)
                            },
                            a = function(t) {
                                return t.slice(0, t.length - 1)
                            },
                            c = e("x" + o);
                        c[""] = i,
                        e("l" + a(o))[u(o)] = c,
                        (n = o, [].concat(Yc(Object.entries(jf)), [["c", "k"]]).reduce((function(t, e) {
                            var r = Jc(e, 2),
                                o = r[0],
                                i = r[1];
                            return n.startsWith(i) ? t.concat(n.replace(i, o)) : t
                        }), [])).forEach((function(t) {
                            ["l", "x"].forEach((function(n) {
                                e(n + a(t))[u(t)] = e(n + o)
                            }))
                        }))
                    })), Object.entries(wf).forEach((function(t) {
                        var n = Jc(t, 2),
                            r = n[1];
                        e(n[0])[""] = r
                    })), [].concat(Yc(Object.keys(bf)), ["c", "y", "w", "j"]).forEach((function(e) {
                        var n = t[e];
                        n[e] = function t(e) {
                            return Object.entries(e).reduce((function(e, n) {
                                var r = Jc(n, 2),
                                    o = r[0],
                                    i = r[1];
                                return e[o] = o ? t(i) : "\u3063" + i, e
                            }), {})
                        }(n)
                    })), delete t.n.n, Object.freeze(JSON.parse(JSON.stringify(t)))
                }
                var Mf = null,
                    zf = pf({
                        wi: "\u3090",
                        we: "\u3091"
                    });
                function Af() {
                    var t = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : "";
                    return !Xc(t) && $c(t, 65, 90)
                }
                function Pf() {
                    var t = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : "";
                    return !Xc(t) && 12540 === t.charCodeAt(0)
                }
                function kf() {
                    var t = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : "";
                    return !Xc(t) && 12539 === t.charCodeAt(0)
                }
                function Lf() {
                    var t = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : "";
                    return !Xc(t) && (!!Pf(t) || $c(t, ef, 12438))
                }
                function Ff() {
                    var t = [];
                    return (arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : "").split("").forEach((function(e) {
                        if (Pf(e) || kf(e))
                            t.push(e);
                        else if (Lf(e)) {
                            var n = e.charCodeAt(0) + 96;
                            t.push(String.fromCharCode(n))
                        } else
                            t.push(e)
                    })), t.join("")
                }
                function Nf() {
                    var t = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : "",
                        e = arguments.length > 1 && void 0 !== arguments[1] ? arguments[1] : {},
                        n = arguments[2],
                        r = void 0;
                    return n ? r = e : n = Tf(r = lf(e)), function() {
                        var t = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : "",
                            e = arguments.length > 1 && void 0 !== arguments[1] ? arguments[1] : {},
                            n = arguments[2];
                        return n || (n = Tf(e)), hf(t.toLowerCase(), n, !e.IMEMode)
                    }(t, r, n).map((function(e) {
                        var n = Jc(e, 3),
                            o = n[0],
                            i = n[1],
                            u = n[2];
                        if (null === u)
                            return t.slice(o);
                        var a = r.IMEMode === Qc.HIRAGANA,
                            c = r.IMEMode === Qc.KATAKANA || [].concat(Yc(t.slice(o, i))).every(Af);
                        return a || !c ? u : Ff(u)
                    })).join("")
                }
                var If = null;
                function Tf() {
                    var t = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : {},
                        e = (null == Mf && (Mf = Sf()), Mf);
                    return e = t.IMEMode ? function(t) {
                        var e = JSON.parse(JSON.stringify(t));
                        return e.n.n = {
                            "": "\u3093"
                        }, e.n[" "] = {
                            "": "\u3093"
                        }, e
                    }(e) : e, e = t.useObsoleteKana ? zf(e) : e, t.customKanaMapping && (null == If && (If = yf(e, t.customKanaMapping)), e = If), e
                }
                var Rf = [];
                function Cf(t) {
                    var e = Object.assign({}, lf(t), {
                            IMEMode: t.IMEMode || !0
                        }),
                        n = Tf(e),
                        r = [].concat(Yc(Object.keys(n)), Yc(Object.keys(n).map((function(t) {
                            return t.toUpperCase()
                        }))));
                    return function(t) {
                        var o = t.target;
                        void 0 !== o.value && "true" !== o.dataset.ignoreComposition && function(t, e, n, r, o) {
                            var i = function() {
                                    var t = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : "",
                                        e = arguments.length > 1 && void 0 !== arguments[1] ? arguments[1] : 0,
                                        n = arguments.length > 2 && void 0 !== arguments[2] ? arguments[2] : [],
                                        r = void 0,
                                        o = void 0,
                                        i = void 0;
                                    if (0 === e && n.includes(t[0])) {
                                        var u = function(t, e) {
                                                return [""].concat(Yc(Df(t, (function(t) {
                                                    return e.includes(t) || !sf(t, /[0-9]/)
                                                }))))
                                            }(t, n),
                                            a = Jc(u, 3);
                                        r = a[0],
                                        o = a[1],
                                        i = a[2]
                                    } else if (e > 0) {
                                        var c = function() {
                                                var t = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : "",
                                                    e = arguments.length > 1 && void 0 !== arguments[1] ? arguments[1] : 0,
                                                    n = Df([].concat(Yc(t.slice(0, e))).reverse(), (function(t) {
                                                        return !sf(t)
                                                    })),
                                                    r = Jc(n, 2),
                                                    o = r[0];
                                                return [r[1].reverse().join(""), o.split("").reverse().join(""), t.slice(e)]
                                            }(t, e),
                                            f = Jc(c, 3);
                                        r = f[0],
                                        o = f[1],
                                        i = f[2]
                                    } else {
                                        var s = Df(t, (function(t) {
                                                return !n.includes(t)
                                            })),
                                            l = Jc(s, 2);
                                        r = l[0];
                                        var h = Df(o = l[1], (function(t) {
                                                return !sf(t)
                                            })),
                                            d = Jc(h, 2);
                                        o = d[0],
                                        i = d[1]
                                    }
                                    return [r, o, i]
                                }(t.value, t.selectionEnd, r),
                                u = Jc(i, 3),
                                a = u[0],
                                c = u[1],
                                f = u[2],
                                s = Nf(c, e, n);
                            if (c !== s) {
                                var l = a.length + s.length,
                                    h = a + s + f;
                                t.value = h,
                                f.length ? setTimeout((function() {
                                    return t.setSelectionRange(l, l)
                                }), 1) : t.setSelectionRange(l, l)
                            } else
                                t.value
                        }(o, e, n, r)
                    }
                }
                function Wf(t) {
                    var e = t.type,
                        n = t.target,
                        r = t.data;
                    /Mac/.test(window.navigator && window.navigator.platform) && ("compositionupdate" === e && sf(r) && (n.dataset.ignoreComposition = "true"), "compositionend" === e && (n.dataset.ignoreComposition = "false"))
                }
                function Df() {
                    for (var t = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : {}, e = arguments.length > 1 && void 0 !== arguments[1] ? arguments[1] : function(t) {
                            return !!t
                        }, n = [], r = t.length, o = 0; r > o && e(t[o], o);)
                        n.push(t[o]),
                        o += 1;
                    return [n.join(""), t.slice(o)]
                }
                var Bf = {
                        input: function(t) {
                            var e = t.target;
                            return console.log("input:", {
                                value: e.value,
                                selectionStart: e.selectionStart,
                                selectionEnd: e.selectionEnd
                            })
                        },
                        compositionstart: function() {
                            return console.log("compositionstart")
                        },
                        compositionupdate: function(t) {
                            var e = t.target;
                            return console.log("compositionupdate", {
                                data: t.data,
                                value: e.value,
                                selectionStart: e.selectionStart,
                                selectionEnd: e.selectionEnd
                            })
                        },
                        compositionend: function() {
                            return console.log("compositionend")
                        }
                    },
                    Kf = function(t) {
                        Object.entries(Bf).forEach((function(e) {
                            var n = Jc(e, 2);
                            return t.addEventListener(n[0], n[1])
                        }))
                    },
                    Vf = function(t) {
                        Object.entries(Bf).forEach((function(e) {
                            var n = Jc(e, 2);
                            return t.removeEventListener(n[0], n[1])
                        }))
                    },
                    Uf = ["TEXTAREA", "INPUT"],
                    Gf = 0,
                    Hf = function() {
                        return Gf += 1, "" + Date.now() + Gf
                    };
                function Jf() {
                    var t = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : "";
                    return !Xc(t) && af.some((function(e) {
                            var n = Jc(e, 2);
                            return $c(t, n[0], n[1])
                        }))
                }
                function Yf() {
                    var t = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : "",
                        e = arguments[1],
                        n = "regexp" === qc(e);
                    return !Xc(t) && [].concat(Yc(t)).every((function(t) {
                            var r = Jf(t);
                            return n ? r || e.test(t) : r
                        }))
                }
                function qf() {
                    return $c(arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : "", nf, 12540)
                }
                function Xf() {
                    var t = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : "";
                    return !Xc(t) && (Lf(t) || qf(t))
                }
                function $f() {
                    var t = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : "";
                    return !Xc(t) && [].concat(Yc(t)).every(Xf)
                }
                function Qf() {
                    var t = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : "";
                    return !Xc(t) && [].concat(Yc(t)).every(Lf)
                }
                function Zf() {
                    var t = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : "";
                    return !Xc(t) && [].concat(Yc(t)).every(qf)
                }
                function ts() {
                    return $c(arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : "", 19968, 40879)
                }
                function es() {
                    var t = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : "";
                    return !Xc(t) && [].concat(Yc(t)).every(ts)
                }
                function ns() {
                    var t = arguments.length > 1 && void 0 !== arguments[1] ? arguments[1] : {
                            passKanji: !0
                        },
                        e = [].concat(Yc(arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : "")),
                        n = !1;
                    return t.passKanji || (n = e.some(es)), (e.some(Qf) || e.some(Zf)) && e.some(Yf) && !n
                }
                var rs = function(t, e) {
                        return Pf(t) && 1 > e
                    },
                    os = function(t, e) {
                        return Pf(t) && e > 0
                    },
                    is = function(t) {
                        return ["\u30f6", "\u30f5"].includes(t)
                    },
                    us = {
                        a: "\u3042",
                        i: "\u3044",
                        u: "\u3046",
                        e: "\u3048",
                        o: "\u3046"
                    };
                function as() {
                    var t = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : "",
                        e = arguments[1],
                        n = arguments[2],
                        r = "";
                    return t.split("").reduce((function(o, i, u) {
                        if (kf(i) || rs(i, u) || is(i))
                            return o.concat(i);
                        if (r && os(i, u)) {
                            var a = e(r).slice(-1);
                            return qf(t[u - 1]) && "o" === a && n ? o.concat("\u304a") : o.concat(us[a])
                        }
                        if (!Pf(i) && qf(i)) {
                            var c = i.charCodeAt(0) + -96,
                                f = String.fromCharCode(c);
                            return r = f, o.concat(f)
                        }
                        return r = "", o.concat(i)
                    }), []).join("")
                }
                var cs = null,
                    fs = {
                        "\u3042": "a",
                        "\u3044": "i",
                        "\u3046": "u",
                        "\u3048": "e",
                        "\u304a": "o",
                        "\u304b": "ka",
                        "\u304d": "ki",
                        "\u304f": "ku",
                        "\u3051": "ke",
                        "\u3053": "ko",
                        "\u3055": "sa",
                        "\u3057": "shi",
                        "\u3059": "su",
                        "\u305b": "se",
                        "\u305d": "so",
                        "\u305f": "ta",
                        "\u3061": "chi",
                        "\u3064": "tsu",
                        "\u3066": "te",
                        "\u3068": "to",
                        "\u306a": "na",
                        "\u306b": "ni",
                        "\u306c": "nu",
                        "\u306d": "ne",
                        "\u306e": "no",
                        "\u306f": "ha",
                        "\u3072": "hi",
                        "\u3075": "fu",
                        "\u3078": "he",
                        "\u307b": "ho",
                        "\u307e": "ma",
                        "\u307f": "mi",
                        "\u3080": "mu",
                        "\u3081": "me",
                        "\u3082": "mo",
                        "\u3089": "ra",
                        "\u308a": "ri",
                        "\u308b": "ru",
                        "\u308c": "re",
                        "\u308d": "ro",
                        "\u3084": "ya",
                        "\u3086": "yu",
                        "\u3088": "yo",
                        "\u308f": "wa",
                        "\u3090": "wi",
                        "\u3091": "we",
                        "\u3092": "wo",
                        "\u3093": "n",
                        "\u304c": "ga",
                        "\u304e": "gi",
                        "\u3050": "gu",
                        "\u3052": "ge",
                        "\u3054": "go",
                        "\u3056": "za",
                        "\u3058": "ji",
                        "\u305a": "zu",
                        "\u305c": "ze",
                        "\u305e": "zo",
                        "\u3060": "da",
                        "\u3062": "ji",
                        "\u3065": "zu",
                        "\u3067": "de",
                        "\u3069": "do",
                        "\u3070": "ba",
                        "\u3073": "bi",
                        "\u3076": "bu",
                        "\u3079": "be",
                        "\u307c": "bo",
                        "\u3071": "pa",
                        "\u3074": "pi",
                        "\u3077": "pu",
                        "\u307a": "pe",
                        "\u307d": "po",
                        "\u3094\u3041": "va",
                        "\u3094\u3043": "vi",
                        "\u3094": "vu",
                        "\u3094\u3047": "ve",
                        "\u3094\u3049": "vo"
                    },
                    ss = {
                        "\u3002": ".",
                        "\u3001": ",",
                        "\uff1a": ":",
                        "\u30fb": "/",
                        "\uff01": "!",
                        "\uff1f": "?",
                        "\u301c": "~",
                        "\u30fc": "-",
                        "\u300c": "\u2018",
                        "\u300d": "\u2019",
                        "\u300e": "\u201c",
                        "\u300f": "\u201d",
                        "\uff3b": "[",
                        "\uff3d": "]",
                        "\uff08": "(",
                        "\uff09": ")",
                        "\uff5b": "{",
                        "\uff5d": "}",
                        "\u3000": " "
                    },
                    ls = ["\u3042", "\u3044", "\u3046", "\u3048", "\u304a", "\u3084", "\u3086", "\u3088"],
                    hs = {
                        "\u3083": "ya",
                        "\u3085": "yu",
                        "\u3087": "yo"
                    },
                    ds = {
                        "\u3043": "yi",
                        "\u3047": "ye"
                    },
                    vs = {
                        "\u3041": "a",
                        "\u3043": "i",
                        "\u3045": "u",
                        "\u3047": "e",
                        "\u3049": "o"
                    },
                    ps = ["\u304d", "\u306b", "\u3072", "\u307f", "\u308a", "\u304e", "\u3073", "\u3074", "\u3094", "\u304f", "\u3075"],
                    ys = {
                        "\u3057": "sh",
                        "\u3061": "ch",
                        "\u3058": "j",
                        "\u3062": "j"
                    },
                    gs = {
                        "\u3063": "",
                        "\u3083": "ya",
                        "\u3085": "yu",
                        "\u3087": "yo",
                        "\u3041": "a",
                        "\u3043": "i",
                        "\u3045": "u",
                        "\u3047": "e",
                        "\u3049": "o"
                    },
                    ms = {
                        b: "b",
                        c: "t",
                        d: "d",
                        f: "f",
                        g: "g",
                        h: "h",
                        j: "j",
                        k: "k",
                        m: "m",
                        p: "p",
                        q: "q",
                        r: "r",
                        s: "s",
                        t: "t",
                        v: "v",
                        w: "w",
                        x: "x",
                        z: "z"
                    };
                function bs() {
                    var t,
                        e,
                        n;
                    return null == cs && (t = df(fs), e = function(e) {
                        return vf(t, e)
                    }, n = function(t, n) {
                        e(t)[""] = n
                    }, Object.entries(ss).forEach((function(t) {
                        var n = Jc(t, 2),
                            r = n[1];
                        e(n[0])[""] = r
                    })), [].concat(Yc(Object.entries(hs)), Yc(Object.entries(vs))).forEach((function(t) {
                        var e = Jc(t, 2);
                        n(e[0], e[1])
                    })), ps.forEach((function(t) {
                        var r = e(t)[""][0];
                        Object.entries(hs).forEach((function(e) {
                            var o = Jc(e, 2);
                            n(t + o[0], r + o[1])
                        })),
                        Object.entries(ds).forEach((function(e) {
                            var o = Jc(e, 2);
                            n(t + o[0], r + o[1])
                        }))
                    })), Object.entries(ys).forEach((function(t) {
                        var e = Jc(t, 2),
                            r = e[0],
                            o = e[1];
                        Object.entries(hs).forEach((function(t) {
                            var e = Jc(t, 2);
                            n(r + e[0], o + e[1][1])
                        })),
                        n(r + "\u3043", o + "yi"),
                        n(r + "\u3047", o + "e")
                    })), t["\u3063"] = function t(e) {
                        return Object.entries(e).reduce((function(e, n) {
                            var r = Jc(n, 2),
                                o = r[0],
                                i = r[1];
                            if (o)
                                e[o] = t(i);
                            else {
                                var u = i.charAt(0);
                                e[o] = Object.keys(ms).includes(u) ? ms[u] + i : i
                            }
                            return e
                        }), {})
                    }(t), Object.entries(gs).forEach((function(t) {
                        var e = Jc(t, 2);
                        n(e[0], e[1])
                    })), ls.forEach((function(t) {
                        n("\u3093" + t, "n'" + e(t)[""])
                    })), cs = Object.freeze(JSON.parse(JSON.stringify(t)))), cs
                }
                function _s() {
                    var t = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : "",
                        e = arguments.length > 1 && void 0 !== arguments[1] ? arguments[1] : {},
                        n = lf(e);
                    return function(t, e) {
                        var n = function(t) {
                            switch (t.romanization) {
                            case Zc.HEPBURN:
                                return bs();
                            default:
                                return {}
                            }
                        }(e);
                        return e.customRomajiMapping && (null == Os && (Os = yf(n, e.customRomajiMapping)), n = Os), hf(as(t, _s, !0), n, !e.IMEMode)
                    }(t, n).map((function(n) {
                        var r = Jc(n, 3),
                            o = r[2];
                        return e.upcaseKatakana && Zf(t.slice(r[0], r[1])) ? o.toUpperCase() : o
                    })).join("")
                }
                var Os = null;
                function js() {
                    var t = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : "";
                    return !Xc(t) && cf.some((function(e) {
                            var n = Jc(e, 2);
                            return $c(t, n[0], n[1])
                        }))
                }
                function Es() {
                    var t = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : "";
                    return !Xc(t) && of.some((function(e) {
                            var n = Jc(e, 2);
                            return $c(t, n[0], n[1])
                        }))
                }
                var ws = function(t) {
                        return " " === t
                    },
                    xs = function(t) {
                        return "\u3000" === t
                    },
                    Ss = function(t) {
                        return /[\uff10-\uff19]/.test(t)
                    },
                    Ms = function(t) {
                        return /[0-9]/.test(t)
                    },
                    zs = "en",
                    As = "ja",
                    Ps = "englishNumeral",
                    ks = "japaneseNumeral",
                    Ls = "englishPunctuation",
                    Fs = "japanesePunctuation",
                    Ns = "kanji",
                    Is = "hiragana",
                    Ts = "katakana",
                    Rs = "space",
                    Cs = "other";
                function Ws(t) {
                    var e = zs,
                        n = As,
                        r = Ps,
                        o = ks,
                        i = Ls,
                        u = Fs,
                        a = Ns,
                        c = Is,
                        f = Ts,
                        s = Rs,
                        l = Cs;
                    if (arguments.length > 1 && void 0 !== arguments[1] && arguments[1])
                        switch (!0) {
                        case Ss(t):
                        case Ms(t):
                            return l;
                        case ws(t):
                            return e;
                        case js(t):
                            return l;
                        case xs(t):
                            return n;
                        case Es(t):
                            return l;
                        case ff(t):
                            return n;
                        case Jf(t):
                            return e;
                        default:
                            return l
                        }
                    else
                        switch (!0) {
                        case xs(t):
                        case ws(t):
                            return s;
                        case Ss(t):
                            return o;
                        case Ms(t):
                            return r;
                        case js(t):
                            return i;
                        case Es(t):
                            return u;
                        case ts(t):
                            return a;
                        case Lf(t):
                            return c;
                        case qf(t):
                            return f;
                        case ff(t):
                            return n;
                        case Jf(t):
                            return e;
                        default:
                            return l
                        }
                }
                function Ds(t) {
                    var e = arguments.length > 1 && void 0 !== arguments[1] ? arguments[1] : {},
                        n = e.compact,
                        r = void 0 !== n && n,
                        o = e.detailed,
                        i = void 0 !== o && o;
                    if (null == t || Xc(t))
                        return [];
                    var u = [].concat(Yc(t)),
                        a = u.shift(),
                        c = Ws(a, r);
                    return u.reduce((function(t, e) {
                        var n = Ws(e, r),
                            o = n === c;
                        c = n;
                        var u = e;
                        return o && (u = (i ? t.pop().value : t.pop()) + u), t.concat(i ? {
                            type: n,
                            value: u
                        } : u)
                    }), [a = i ? {
                        type: c,
                        value: a
                    } : a])
                }
                var Bs = function(t, e) {
                        return e && !$f(t[0])
                    },
                    Ks = function(t, e) {
                        return !e && !$f(t[t.length - 1])
                    },
                    Vs = function(t, e) {
                        return e && ![].concat(Yc(e)).some(es) || !e && $f(t)
                    };
                t.bind = function() {
                    var t = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : {},
                        e = arguments.length > 1 && void 0 !== arguments[1] ? arguments[1] : {},
                        n = arguments.length > 2 && void 0 !== arguments[2] && arguments[2];
                    if (!Uf.includes(t.nodeName))
                        throw Error("Element provided to Wanakana bind() was not a valid input or textarea element.\n Received: (" + JSON.stringify(t) + ")");
                    var r = Cf(e),
                        o = Hf();
                    t.setAttribute("data-wanakana-id", o),
                    t.setAttribute("lang", "ja"),
                    t.setAttribute("autoCapitalize", "none"),
                    t.setAttribute("autoCorrect", "off"),
                    t.setAttribute("autoComplete", "off"),
                    t.setAttribute("spellCheck", "false"),
                    t.addEventListener("input", r),
                    t.addEventListener("compositionupdate", Wf),
                    t.addEventListener("compositionend", Wf),
                    function(t, e, n) {
                        Rf = Rf.concat({
                            id: t,
                            inputHandler: e,
                            compositionHandler: n
                        })
                    }(o, r, Wf),
                    !0 === n && Kf(t)
                },
                t.unbind = function(t) {
                    var e,
                        n = arguments.length > 1 && void 0 !== arguments[1] && arguments[1],
                        r = (e = t) && Rf.find((function(t) {
                            return t.id === e.getAttribute("data-wanakana-id")
                        }));
                    if (null == r)
                        throw Error("Element provided to Wanakana unbind() had no listener registered.\n Received: " + JSON.stringify(t));
                    var o,
                        i = r.inputHandler,
                        u = r.compositionHandler;
                    t.removeAttribute("data-wanakana-id"),
                    t.removeAttribute("data-ignore-composition"),
                    t.removeEventListener("input", i),
                    t.removeEventListener("compositionstart", u),
                    t.removeEventListener("compositionupdate", u),
                    t.removeEventListener("compositionend", u),
                    o = r.id,
                    Rf = Rf.filter((function(t) {
                        return t.id !== o
                    })),
                    !0 === n && Vf(t)
                },
                t.isRomaji = Yf,
                t.isJapanese = sf,
                t.isKana = $f,
                t.isHiragana = Qf,
                t.isKatakana = Zf,
                t.isMixed = ns,
                t.isKanji = es,
                t.toRomaji = _s,
                t.toKana = Nf,
                t.toHiragana = function() {
                    var t = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : "",
                        e = lf(arguments.length > 1 && void 0 !== arguments[1] ? arguments[1] : {});
                    return e.passRomaji ? as(t, _s) : ns(t, {
                        passKanji: !0
                    }) ? Nf(as(t, _s).toLowerCase(), e) : Yf(t) || js(t) ? Nf(t.toLowerCase(), e) : as(t, _s)
                },
                t.toKatakana = function() {
                    var t = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : "",
                        e = lf(arguments.length > 1 && void 0 !== arguments[1] ? arguments[1] : {});
                    return e.passRomaji ? Ff(t) : ns(t) || Yf(t) || js(t) ? Ff(Nf(t.toLowerCase(), e)) : Ff(t)
                },
                t.stripOkurigana = function() {
                    var t = arguments.length > 0 && void 0 !== arguments[0] ? arguments[0] : "",
                        e = arguments.length > 1 && void 0 !== arguments[1] ? arguments[1] : {},
                        n = e.leading,
                        r = void 0 !== n && n,
                        o = e.matchKanji,
                        i = void 0 === o ? "" : o;
                    if (!sf(t) || Bs(t, r) || Ks(t, r) || Vs(t, i))
                        return t;
                    var u = i || t,
                        a = RegExp(r ? "^" + Ds(u).shift() : Ds(u).pop() + "$");
                    return t.replace(a, "")
                },
                t.tokenize = Ds,
                t.VERSION = "4.0.2",
                t.TO_KANA_METHODS = Qc,
                t.ROMANIZATIONS = Zc,
                Object.defineProperty(t, "__esModule", {
                    value: !0
                })
            },
            "object" == a(e) && "undefined" != typeof t ? u(e) : (o = [e], void 0 === (i = "function" === typeof (r = u) ? r.apply(e, o) : r) || (t.exports = i))
        }).call(this, n(76))
    },
    307: function(t, e, n) {
        "use strict";
        var r = {
            onyomi: "on",
            kunyomi: "kun",
            nanori: "nanori"
        };
        e.a = function(t) {
            return t.kana ? t.kana : function(t) {
                var e = null === t || void 0 === t ? void 0 : t.emph;
                return (null === t || void 0 === t ? void 0 : t[null === r || void 0 === r ? void 0 : r[e]]) ? t[r[t.emph]] : []
            }(t)
        }
    },
    53: function(t, e, n) {
        "use strict";
        var r = n(26);
        e.a = function(t) {
            return t.split("").map(r.toHiragana).join("")
        }
    },
    76: function(t, e) {
        function n(t) {
            return (n = "function" === typeof Symbol && "symbol" === typeof Symbol.iterator ? function(t) {
                return typeof t
            } : function(t) {
                return t && "function" === typeof Symbol && t.constructor === Symbol && t !== Symbol.prototype ? "symbol" : typeof t
            })(t)
        }
        var r;
        r = function() {
            return this
        }();
        try {
            r = r || new Function("return this")()
        } catch (o) {
            "object" === ("undefined" === typeof window ? "undefined" : n(window)) && (r = window)
        }
        t.exports = r
    },
    851: function(t, e, n) {
        "use strict";
        var r = n(26),
            o = {
                "\u3062": "di",
                "\u3065": "du",
                "\u3062\u3083": "dya",
                "\u3062\u3085": "dyu",
                "\u3062\u3087": "dyo",
                "\u3075": "hu"
            };
        e.a = function(t) {
            return Object(r.toRomaji)(t, {
                customRomajiMapping: o
            })
        }
    },
    976: function(t, e, n) {
        "use strict";
        n.r(e);
        var r = function(t) {
            return t ? t.voc || t.kan || t.rad : null
        };
        function o(t, e, n) {
            return r(n) === e || null
        }
        var i = n(26),
            u = n(53),
            a = ["\u3083", "\u3085", "\u3087", "\u30e3", "\u30e5", "\u30e7"],
            c = ["\u3063", "\u30c3"];
        function f(t, e) {
            var n = -1 !== c.indexOf(t[e]),
                r = -1 !== a.indexOf(t[e + 1]);
            return {
                hasSmallKana: n || r,
                isSurrounded: n && -1 !== a.indexOf(t[e + 2])
            }
        }
        function s(t) {
            return function(t) {
                for (var e = [], n = 0; n < t.length; n += 1) {
                    var r = f(t, n),
                        o = r.hasSmallKana;
                    r.isSurrounded ? (e.push(t[n] + t[n + 1] + t[n + 2]), n += 2) : o ? (e.push(t[n] + t[n + 1]), n += 1) : e.push(t[n])
                }
                return e
            }(t).map((function(t) {
                return "\u3093" === t || "\u30f3" === t ? "nn" : "\u30fc" === t ? "-" : Object(i.toRomaji)(t)
            })).join("")
        }
        var l = n(307),
            h = {
                "\u304a": "\u3046",
                "\u3046": "\u304a",
                "\u3048": "\u3044"
            };
        var d = function(t, e, n) {
            return t.length === e.length && t.split("").every((function(t, r) {
                    return function(t, e, n, r) {
                        if (n === t[r])
                            return !0;
                        var o = h[n] === t[r],
                            i = "\u30fc" === e[r],
                            u = -1 === Object.keys(h).indexOf(e[r - 1]);
                        return o && i && u
                    }(e, n, t, r)
                }))
        };
        function v(t, e, n) {
            if ("reading" !== t || !n)
                return null;
            var r = (Object(l.a)(n) || []).find((function(t) {
                return -1 !== t.indexOf("\u30fc") && d(function(t) {
                        for (var e, n, r = "", o = 0; o < t.length; o += 1)
                            "\u30fc" === t[o] ? r += Object(u.a)((e = t[o - 1], n = void 0, (n = Object(i.toRomaji)(e))[n.length - 1])) : r += Object(u.a)(t[o]);
                        return r
                    }(t), Object(u.a)(e), t)
            }));
            return r ? "Try typing \u201c".concat(s(r), "\u201d to get that long \u30fc.") : null
        }
        var p = n(851),
            y = function(t) {
                return t.split("").map((function(t) {
                    return "\u3063" === t ? "*" : Object(p.a)(t)
                }))
            },
            g = function(t) {
                return t.startsWith("n") && t.length > 1
            },
            m = function(t, e, n) {
                return t.every((function(t, r) {
                    return function(t, e, n) {
                        return t === e && !g(t) || "n".concat(t) === e || n && t === "n".concat(e) || g(t) && "n" === e
                    }(t, e[r], n)
                }))
            },
            b = function(t, e) {
                return t.every((function(t, n) {
                    return t === e[n]
                }))
            },
            _ = function(t, e) {
                return !b(t, e) && (t.length === e.length ? m(t, e) : (n = e, r = t.slice(), n.forEach((function(t, e) {
                        "n" === t && r[e] && r[e].startsWith("n") && "n" !== r[e] && e !== r.length && r.splice(e, 0, "n")
                    })), r.length === n.length && (b(r, n) || m(r, n, !0))));
                var n,
                    r
            };
        function O(t, e, n) {
            var r = function(t) {
                return t && Object.prototype.hasOwnProperty.call(t, "kana") ? t.kana.filter((function(t) {
                    return -1 !== t.indexOf("\u3093")
                })) : []
            }(n);
            if ("reading" === t && (null === n || void 0 === n ? void 0 : n.kana) && !n.kana.find((function(t) {
                return t === e
            })) && r.length > 0) {
                var o = y(e),
                    i = r.map(y),
                    u = r[i.findIndex((function(t) {
                        return _(o, t)
                    }))];
                if (u)
                    return "Don\u2019t forget that \u3093 is typed as \u201cnn\u201d. Try typing \u201c".concat(s(u), "\u201d.")
            }
            return null
        }
        var j = {
                meaning: "en",
                reading: "ja"
            },
            E = function(t, e) {
                return "reading" === e ? Object(i.toHiragana)(t.trim()) : t.trim().toLowerCase()
            };
        function w(t, e, n) {
            var r = (null === n || void 0 === n ? void 0 : n.voc) || "",
                o = null === n || void 0 === n ? void 0 : n.kanji;
            return o && 0 !== o.length && 1 === r.length && -1 !== o.reduce((function(e, n) {
                var r = (n[j[t]] || "").split(",");
                return e.concat(r.map((function(e) {
                    return E(e, t)
                })))
            }), []).indexOf(E(e, t)) ? "Oops, we want the vocabulary ".concat(t, ", not the kanji ").concat(t, ".") : null
        }
        var x = function(t) {
                return 1 === t.length ? t[0] : 2 === t.length ? "".concat(t[0], " and ").concat(t[1]) : "".concat(t.slice(0, t.length - 1).join(", "), " and ").concat(t[t.length - 1])
            },
            S = {
                "\u3083": "\u3084",
                "\u3085": "\u3086",
                "\u3087": "\u3088",
                "\u30e3": "\u30e4",
                "\u30e5": "\u30e6",
                "\u30e7": "\u30e8"
            },
            M = "typo",
            z = "mistake";
        function A(t, e) {
            return function(t, e) {
                return t === S[e] && e
            }(t, e) ? M : t === e || z
        }
        function P(t, e) {
            return t.length === e.length && function(t, e) {
                    for (var n = [], r = 0; r < t.length; r += 1) {
                        var o = t[r],
                            i = e[r],
                            u = A(o, i);
                        if (u === M && n.push({
                            expectedChar: i,
                            expectedAnswer: e
                        }), u === z)
                            return !1
                    }
                    return n
                }(t, e)
        }
        function k(t, e, n) {
            if ("reading" !== t)
                return null;
            var r = Object(l.a)(n).map((function(t) {
                    return P(e, t)
                })).filter((function(t) {
                    return t
                }))[0],
                o = r && r.length > 0 && r.map((function(t) {
                    return t.expectedChar
                }));
            return r && r.length > 0 ? "Watch out for the small ".concat(x(o), ". Try typing \u201c").concat(s(r[0].expectedAnswer), "\u201d for this one.") : null
        }
        function L(t, e, n, r) {
            if (!r.passed && "meaning" === t && n.en && (null === n || void 0 === n ? void 0 : n.en.some((function(t) {
                return t.toLowerCase().startsWith("to ")
            }))) && !e.toLowerCase().startsWith("to ")) {
                var o = n.en.find((function(t) {
                    return function(t, e) {
                        var n = t.toLowerCase();
                        return n.startsWith("to ") && n.replace("to ", "") === e.toLowerCase()
                    }(t, e)
                }));
                if (o && n.voc) {
                    var u = n.voc.replace(Object(i.stripOkurigana)(n.voc), "");
                    return "Almost! It ends in ".concat(u[u.length - 1], ", so it\u2019s a verb. Please type \u201c").concat(o, "\u201d.")
                }
            }
            return null
        }
        var F = /nn/gi;
        function N(t, e, n, r) {
            if ("meaning" === t && !r.passed && Object(l.a)(n).some((function(t) {
                return function(t, e) {
                    return Object(i.toHiragana)(t) === Object(i.toHiragana)(e.replace(F, "n"))
                }(t, e)
            })))
                return "Oops, we want the meaning, not the reading.";
            return null
        }
        function I(t, e, n) {
            var r = function(t, e) {
                return ((null === e || void 0 === e ? void 0 : e["auxiliary_".concat(t, "s")]) || []).filter((function(t) {
                    return "warn" === t.type
                }))
            }(t, n).find((function(n) {
                return n[t].toLowerCase() === e.toLowerCase()
            }));
            return r ? r.message : null
        }
        var T = n(161);
        function R(t, e) {
            var n = Object.keys(t);
            if (Object.getOwnPropertySymbols) {
                var r = Object.getOwnPropertySymbols(t);
                e && (r = r.filter((function(e) {
                    return Object.getOwnPropertyDescriptor(t, e).enumerable
                }))),
                n.push.apply(n, r)
            }
            return n
        }
        function C(t) {
            for (var e = 1; e < arguments.length; e++) {
                var n = null != arguments[e] ? arguments[e] : {};
                e % 2 ? R(Object(n), !0).forEach((function(e) {
                    W(t, e, n[e])
                })) : Object.getOwnPropertyDescriptors ? Object.defineProperties(t, Object.getOwnPropertyDescriptors(n)) : R(Object(n)).forEach((function(e) {
                    Object.defineProperty(t, e, Object.getOwnPropertyDescriptor(n, e))
                }))
            }
            return t
        }
        function W(t, e, n) {
            return e in t ? Object.defineProperty(t, e, {
                value: n,
                enumerable: !0,
                configurable: !0,
                writable: !0
            }) : t[e] = n, t
        }
        function D(t, e, n, r, o) {
            for (var i = 0; i < t.length; i += 1) {
                var u = t[i](e, n, r, o);
                if (u)
                    return u
            }
            return null
        }
        function B(t, e) {
            return t.exception && (t.exception = "WaniKani is looking for the ".concat(Object(T.a)(e.emph), " reading.")), t
        }
        var K = function(t, e, n, r) {
                return (t.passed && !t.accurate || !t.passed) && !function(t, e, n) {
                        return !("meaning" !== t || !(null === n || void 0 === n ? void 0 : n.syn)) && n.syn.find((function(t) {
                                return t.trim().toLowerCase() === e.trim().toLowerCase()
                            }))
                    }(e, n, r)
            },
            V = function(t) {
                var e = arguments.length > 1 && void 0 !== arguments[1] ? arguments[1] : [w, N, I, o, v, L, k, O];
                return C(C({}, t), {}, {
                    evaluate: function(n, r) {
                        var o = arguments.length > 2 && void 0 !== arguments[2] ? arguments[2] : $.jStorage.get("currentItem"),
                            i = t.evaluate(n, r, o);
                        if (K(i, n, r, o)) {
                            var u = D(e, n, r, o, i);
                            if (u)
                                return {
                                    exception: u
                                }
                        }
                        return B(i, o)
                    }
                })
            };
        window.enhanceAnswerChecker = V
    }
});
//# sourceMappingURL=answer_checker-115c5ea143b59ee34378.js.map

