"""This is an automatically generated stub for `unregistered.capnp`."""
import os

import capnp  # type: ignore

capnp.remove_import_hook()
here = os.path.dirname(os.path.abspath(__file__))
module_file = os.path.abspath(os.path.join(here, "unregistered.capnp"))
module = capnp.load(module_file)  # pylint: disable=no-member
Unregistered = module.Unregistered
UnregisteredBuilder = Unregistered
UnregisteredReader = Unregistered
