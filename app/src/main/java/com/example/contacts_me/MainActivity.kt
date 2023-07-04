package com.example.contacts_me

import android.app.AlertDialog
import android.os.Bundle
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.BaseAdapter
import android.widget.EditText
import android.widget.ListView
import android.widget.TextView
import androidx.appcompat.app.AppCompatActivity


class MainActivity : AppCompatActivity() {
    private val contacts = ArrayList<Contact>()
    private lateinit var listView: ListView
    private lateinit var adapter: ContactAdapter

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)

        listView = findViewById(R.id.list_contacts)
        adapter = ContactAdapter()
        listView.adapter = adapter

        val addContactButton: View = findViewById(R.id.btn_add_contact)
        addContactButton.setOnClickListener {
            showAddContactDialog()
        }

        listView.setOnItemLongClickListener { _, _, position, _ ->
            showDeleteContactDialog(position)
            true
        }
    }

    private fun showAddContactDialog() {
        val dialogView = LayoutInflater.from(this).inflate(R.layout.dialog_add_contact, null)
        val nameEditText: EditText = dialogView.findViewById(R.id.edit_text_name)
        val phoneEditText: EditText = dialogView.findViewById(R.id.edit_text_phone)

        AlertDialog.Builder(this)
            .setTitle("Add Contact")
            .setView(dialogView)
            .setPositiveButton("Add") { _, _ ->
                val name = nameEditText.text.toString().trim()
                val phoneNumber = phoneEditText.text.toString().trim()

                if (name.isNotEmpty() && phoneNumber.isNotEmpty()) {
                    val contact = Contact(name, phoneNumber)
                    contacts.add(contact)
                    adapter.notifyDataSetChanged()
                }
            }
            .setNegativeButton("Cancel", null)
            .show()
    }

    private fun showDeleteContactDialog(position: Int) {
        AlertDialog.Builder(this)
            .setTitle("Delete Contact")
            .setMessage("Are you sure you want to delete this contact?")
            .setPositiveButton("Delete") { _, _ ->
                contacts.removeAt(position)
                adapter.notifyDataSetChanged()
            }
            .setNegativeButton("Cancel", null)
            .show()
    }

    inner class ContactAdapter : BaseAdapter() {
        override fun getCount(): Int {
            return contacts.size
        }

        override fun getItem(position: Int): Contact {
            return contacts[position]
        }

        override fun getItemId(position: Int): Long {
            return position.toLong()
        }

        override fun getView(position: Int, convertView: View?, parent: ViewGroup?): View {
            val view: View = convertView ?: layoutInflater.inflate(R.layout.contact_item, parent, false)

            val nameTextView: TextView = view.findViewById(R.id.text_view_name)
            val phoneTextView: TextView = view.findViewById(R.id.text_view_phone)

            val contact = getItem(position)
            nameTextView.text = contact.name
            phoneTextView.text = contact.phoneNumber

            return view
        }
    }
}
