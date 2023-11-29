(function () {
	if (document) {
		var g = document.currentScript,
			h = document.location.protocol + '//' + g.getAttribute('data-api'),
			k = Date.now().toString(36) + Math.random().toString(36).substr(2),
			d = !0,
			e = 0,
			f = 0,
			b = new XMLHttpRequest(),
			m = history.pushState,
			n = history.replaceState,
			l = function () {
				d = !1;
				k = Date.now().toString(36) + Math.random().toString(36).substr(2);
				f = e = 0;
			},
			a = function (c) {
				var p = {
					b: k,
					u: location.href,
					r: document.referrer,
					d: Intl.DateTimeFormat().resolvedOptions().timeZone,
					p: d,
					t: document.title,
					w: self.screen.width,
					h: self.screen.height,
					e: c,
					m:
						'pagehide' === c || 'hidden' === c || 'unload' === c
							? self.performance.now() - e
							: void 0,
				};
				navigator.sendBeacon(h + '/event/hit', JSON.stringify(p));
				'unload' === c && l();
			};
		'onpagehide' in self
			? document.addEventListener(
					'pagehide',
					function () {
						a('pagehide');
					},
					{ capture: !0 }
			  )
			: document.addEventListener(
					'unload',
					function () {
						a('unload');
					},
					{ capture: !0 }
			  );
		document.addEventListener(
			'visibilitychange',
			function () {
				'hidden' === document.visibilityState
					? ((f = self.performance.now()), a('hidden'))
					: (e += self.performance.now() - f);
			},
			{ capture: !0 }
		);
		b.open('GET', h + '/event/ping');
		b.setRequestHeader('Content-Type', 'text/plain');
		b.addEventListener(
			'load',
			function () {
				'1' === b.responseText && (d = !1);
				a('load');
				g.getAttribute('data-hash')
					? document.addEventListener(
							'hashchange',
							function () {
								a('load');
							},
							{ capture: !0 }
					  )
					: ((history.pushState = function () {
							a('unload');
							m.apply(history, arguments);
							a('load');
					  }),
					  (history.replaceState = function () {
							a('unload');
							n.apply(history, arguments);
							a('load');
					  }),
					  window.addEventListener(
							'popstate',
							function () {
								l();
								a('load');
							},
							{ capture: !0 }
					  ));
			},
			{ capture: !0, once: !0 }
		);
		b.send();
	}
})();
