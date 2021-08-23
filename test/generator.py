# This file is part of arduino-cli.
#
# Copyright 2020 ARDUINO SA (http://www.arduino.cc/)
#
# This software is released under the GNU General Public License version 3,
# which covers the main part of arduino-cli.
# The terms of this license can be found at:
# https://www.gnu.org/licenses/gpl-3.0.en.html
#
# You can be released from the requirements of the above licenses by purchasing
# a commercial license. Buying such a license is mandatory if you want to modify or
# otherwise use the software for commercial activities involving the Arduino
# software without disclosing the source code of your own applications. To purchase
# a commercial license, send an email to license@arduino.cc.

import tempfile
import hashlib
import pytest
import pprint
import os
from pathlib import Path


def generate_build_dir(sketch_path):
    sketch_path_md5 = hashlib.md5(bytes(sketch_path)).hexdigest().upper()
    build_dir = Path(tempfile.gettempdir(), f"arduino-sketch-{sketch_path_md5}")
    build_dir.mkdir(parents=True, exist_ok=True)
    return build_dir.resolve()


indexes = [
    "https://adafruit.github.io/arduino-board-index/package_adafruit_index.json",
    "https://dl.espressif.com/dl/package_esp32_index.json",
    "http://arduino.esp8266.com/stable/package_esp8266com_index.json",
]

cores_to_install = [
    "arduino:avr@1.8.3",
    "adafruit:avr@1.4.13",
    "arduino:samd@1.8.11",
    "esp32:esp32@1.0.6",
    "esp8266:esp8266@3.0.2",
]

testdata = [
    ("arduino:avr:uno", "/dev/ttyACM0", ""),
    ("arduino:avr:uno", "", "usbasp"),
    ("arduino:avr:uno", "/dev/ttyACM0", "avrisp"),
    ("arduino:avr:uno", "/dev/ttyACM0", "arduinoasisp"),
    ("arduino:avr:nano", "/dev/ttyACM0", ""),
    ("arduino:avr:nano", "", "usbasp"),
    ("arduino:avr:nano", "/dev/ttyACM0", "avrisp"),
    ("arduino:avr:nano", "/dev/ttyACM0", "arduinoasisp"),
    ("arduino:avr:nano:cpu=atmega328old", "/dev/ttyACM0", ""),
    ("arduino:avr:nano:cpu=atmega328old", "", "usbasp"),
    ("arduino:avr:nano:cpu=atmega328old", "/dev/ttyACM0", "avrisp"),
    ("arduino:avr:nano:cpu=atmega328old", "/dev/ttyACM0", "arduinoasisp"),
    ("arduino:avr:mega", "/dev/ttyACM0", ""),
    ("arduino:avr:mega:cpu=atmega1280", "/dev/ttyACM0", ""),
    ("arduino:avr:diecimila", "/dev/ttyACM0", ""),
    ("arduino:avr:leonardo", "/dev/ttyACM0", ""),
    ("arduino:avr:leonardo", "/dev/ttyACM999", ""),
    ("arduino:avr:micro", "/dev/ttyACM0", ""),
    ("arduino:avr:micro", "/dev/ttyACM999", ""),
    ("arduino:avr:circuitplay32u4cat", "/dev/ttyACM0", ""),
    ("arduino:avr:circuitplay32u4cat", "/dev/ttyACM999", ""),
    ("arduino:avr:gemma", "/dev/ttyACM0", "usbGemma"),
    ("arduino:avr:gemma", "", "usbGemma"),
    ("arduino:avr:unowifi", "/dev/ttyACM0", ""),
    ("arduino:avr:yun", "/dev/ttyACM0", ""),
    ("arduino:avr:yun", "/dev/ttyACM999", ""),
    ("adafruit:avr:circuitplay32u4cat", "/dev/ttyACM0", ""),
    ("adafruit:avr:circuitplay32u4cat", "/dev/ttyACM999", ""),
    ("adafruit:avr:flora8", "/dev/ttyACM0", ""),
    ("adafruit:avr:flora8", "/dev/ttyACM999", ""),
    ("adafruit:avr:gemma", "/dev/ttyACM0", "usbGemma"),
    ("adafruit:avr:gemma", "", "usbGemma"),
    ("adafruit:avr:itsybitsy32u4_3V", "/dev/ttyACM0", ""),
    ("adafruit:avr:itsybitsy32u4_3V", "/dev/ttyACM999", ""),
    ("adafruit:avr:itsybitsy32u4_5V", "/dev/ttyACM0", ""),
    ("adafruit:avr:itsybitsy32u4_5V", "/dev/ttyACM999", ""),
    ("adafruit:avr:metro", "/dev/ttyACM0", ""),
    ("adafruit:avr:trinket3", "", "usbasp"),
    ("adafruit:avr:trinket3", "/dev/ttyACM0", "avrisp"),
    ("adafruit:avr:trinket3", "/dev/ttyACM0", "arduinoasisp"),
    ("arduino:samd:arduino_zero_edbg", "", ""),
    ("arduino:samd:adafruit_circuitplayground_m0", "/dev/ttyACM0", ""),
    ("arduino:samd:adafruit_circuitplayground_m0", "/dev/ttyACM999", ""),
    ("arduino:samd:mkrfox1200", "/dev/ttyACM0", ""),
    ("arduino:samd:mkrfox1200", "/dev/ttyACM999", ""),
    ("arduino:samd:mkrgsm1400", "/dev/ttyACM0", ""),
    ("arduino:samd:mkrgsm1400", "/dev/ttyACM999", ""),
    ("arduino:samd:mkrvidor4000", "/dev/ttyACM0", ""),
    ("arduino:samd:mkrvidor4000", "/dev/ttyACM999", ""),
    ("arduino:samd:mkrwan1310", "/dev/ttyACM0", ""),
    ("arduino:samd:mkrwan1310", "/dev/ttyACM999", ""),
    ("arduino:samd:mkrwifi1010", "/dev/ttyACM0", ""),
    ("arduino:samd:mkrwifi1010", "/dev/ttyACM999", ""),
    ("arduino:samd:mkr1000", "/dev/ttyACM0", ""),
    ("arduino:samd:mkr1000", "/dev/ttyACM999", ""),
    ("arduino:samd:mkrzero", "/dev/ttyACM0", ""),
    ("arduino:samd:mkrzero", "/dev/ttyACM999", ""),
    ("arduino:samd:nano_33_iot", "/dev/ttyACM0", ""),
    ("arduino:samd:nano_33_iot", "/dev/ttyACM999", ""),
    ("arduino:samd:arduino_zero_native", "/dev/ttyACM0", ""),
    ("arduino:samd:arduino_zero_native", "/dev/ttyACM999", ""),
    ("esp32:esp32:esp32", "/dev/ttyACM0", ""),
    (
        "esp32:esp32:esp32:PSRAM=enabled,PartitionScheme=no_ota,CPUFreq=80,FlashMode=dio,FlashFreq=40,FlashSize=8M,UploadSpeed=230400,DebugLevel=info",  # noqa: E501
        "/dev/ttyACM0",
        "",
    ),
    ("esp32:esp32:esp32thing", "/dev/ttyACM0", ""),
    ("esp8266:esp8266:generic", "/dev/ttyACM0", ""),
    (
        "esp8266:esp8266:generic:xtal=160,vt=heap,mmu=3216,ResetMethod=nodtr_nosync,CrystalFreq=40,FlashFreq=20,eesz=2M,baud=57600",  # noqa: E501
        "/dev/ttyACM0",
        "",
    ),
]


@pytest.mark.parametrize("fqbn, upload_port, programmer", testdata)
def test_generate_upload_sketch_golden(
    run_command,
    session_data_dir,
    downloads_dir,
    fqbn,
    upload_port,
    programmer,
):
    env = {
        "ARDUINO_DATA_DIR": session_data_dir,
        "ARDUINO_DOWNLOADS_DIR": downloads_dir,
        "ARDUINO_SKETCHBOOK_DIR": session_data_dir,
    }

    # Install everything just once
    if not os.path.isdir(session_data_dir + "/packages"):
        assert run_command("config init --overwrite", custom_env=env)
        for package_index in indexes:
            assert run_command(f"config add board_manager.additional_urls {package_index}", custom_env=env)
        assert run_command("update", custom_env=env)

        for d in cores_to_install:
            assert run_command(f"core install {d}", custom_env=env)

    # Create a sketch
    sketch_name = "TestSketchForUpload"
    sketch_path = Path(session_data_dir, sketch_name)
    assert run_command(f'sketch new "{sketch_path}"', custom_env=env)

    # Fake compilation, we just need the folder to exist
    build_dir = generate_build_dir(sketch_path)

    def run_cli(os):
        env["ARDUINO_CLI_FORCE_OS_SUFFIX"] = os
        if programmer != "":
            programmer_arg = "-P " + programmer
        else:
            programmer_arg = ""
        if upload_port != "":
            port_arg = "-p " + upload_port
        else:
            port_arg = ""
        res = run_command(f'upload {port_arg} {programmer_arg} -b {fqbn} "{sketch_path}" --dry-run -v', custom_env=env)
        assert res.ok
        output = res.stdout.replace("\\", "/")
        output = output.replace("{", "{{")
        output = output.replace("}", "}}")
        output = output.replace(session_data_dir, "{data_dir}")
        output = output.replace(str(build_dir), "{build_dir}")
        output = output.replace(sketch_name, "{sketch_name}")
        output = output.replace("\r", "")
        return output

    linux_output = run_cli("linux")
    macosx_output = run_cli("macosx")
    windows_output = run_cli("windows")

    if linux_output != macosx_output or linux_output != windows_output:
        out = (
            fqbn,
            upload_port,
            programmer,
            {
                "linux": linux_output,
                "win32": windows_output,  # python uses "win32" instead of "windows"
                "darwin": macosx_output,  # python uses "darwin" instead of "macosx"
            },
        )
    else:
        out = (
            fqbn,
            upload_port,
            programmer,
            linux_output,
        )
    with open("test_upload_mock_generated.txt", "a") as outfile:
        pprint.pprint(out, stream=outfile, indent=4)
        outfile.write(",\n")
