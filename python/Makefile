setup: clean install_uv
	mkdir wamp_msgs_capnp/schemas/
	cp ../capnproto/* wamp_msgs_capnp/schemas/
	sed -i '/^using Go/ s/^/#/' wamp_msgs_capnp/schemas/*.capnp
	sed -i '/^\$$Go\./ s/^/#/' wamp_msgs_capnp/schemas/*.capnp
	uv venv .venv --python 3.12
	uv pip install .[test]

clean:
	rm -rf wamp_msgs_capnp/schemas wampproto_messages_capnproto.egg-info .venv build

format:
	. .venv/bin/activate; ruff format .

lint:
	. .venv/bin/activate; ruff check .

install_uv:
	@if ! command -v uv >/dev/null 2>&1; then \
  		curl -LsSf https://astral.sh/uv/install.sh | sh; \
  	fi
