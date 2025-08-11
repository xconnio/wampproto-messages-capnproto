"""This is an automatically generated stub for `unregister.capnp`."""
import os

import capnp  # type: ignore

capnp.remove_import_hook()
here = os.path.dirname(os.path.abspath(__file__))
module_file = os.path.abspath(os.path.join(here, "unregister.capnp"))
module = capnp.load(module_file)  # pylint: disable=no-member
Unregister = module.Unregister
UnregisterBuilder = Unregister
UnregisterReader = Unregister
