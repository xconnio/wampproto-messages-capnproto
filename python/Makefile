setup: clean
	mkdir wamp_msgs_capnp/schemas/
	cp ../capnproto/* wamp_msgs_capnp/schemas/
	sed -i '/^using Go/ s/^/#/' wamp_msgs_capnp/schemas/*.capnp
	sed -i '/^\$$Go\./ s/^/#/' wamp_msgs_capnp/schemas/*.capnp
	uv venv .venv --python 3.12
	uv pip install .

clean:
	rm -rf wamp_msgs_capnp/schemas wampproto_messages_capnproto.egg-info .venv build
